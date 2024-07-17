# common logic across linux and darwin
# (0.1.48)


# 这个函数 init_vars 用于初始化一些变量，根据系统的架构类型（GOARCH）、调试标志（CGO_CFLAGS）和操作系统类型（uname -s），来设置不同的变量值。

# 根据 GOARCH 的值来设置 ARCH 变量，支持 amd64 和 arm64，对于其他架构类型，则通过 uname -m 命令获取当前机器的架构类型，并通过 sed 命令将 aarch64 替换为 arm64。

# 设置 LLAMACPP_DIR 变量为 ../llama.cpp，用于指定一个目录路径。
# 设置 CMAKE_DEFS 变量的初始值为空字符串。
# 设置 CMAKE_TARGETS 变量的初始值为 "--target ollama_llama_server"，用于指定 CMake 构建的目标。
# 根据 CGO_CFLAGS 是否包含 -g 标志来判断是否需要生成调试信息，如果包含，则将一些调试和优化相关的 CMake 定义追加到 CMAKE_DEFS 中，否则，将一些优化相关的 CMake 定义追加到 CMAKE_DEFS 中。

# 根据操作系统的类型（uname -s）来设置不同的变量值，对于 Darwin 系统，设置 LIB_EXT 为 dylib，WHOLE_ARCHIVE 为 -Wl,-force_load，NO_WHOLE_ARCHIVE 为空字符串，GCC_ARCH 为 -arch ${ARCH}；
# 对于 Linux 系统，设置 LIB_EXT 为 so，WHOLE_ARCHIVE 为 -Wl,--whole-archive，NO_WHOLE_ARCHIVE 为 -Wl,--no-whole-archive，GCC_ARCH 为空字符串；对于其他系统类型，不做任何处理。

# 如果 CMAKE_CUDA_ARCHITECTURES 变量为空，则将其设置为 "50;52;61;70;75;80"，用于指定 CUDA 的架构版本。
init_vars() {
    case "${GOARCH}" in
    "amd64")
        ARCH="x86_64"
        ;;
    "arm64")
        ARCH="arm64"
        ;;
    *)
        ARCH=$(uname -m | sed -e "s/aarch64/arm64/g")
    esac

    LLAMACPP_DIR=../llama.cpp
    CMAKE_DEFS=""
    CMAKE_TARGETS="--target ollama_llama_server"
    if echo "${CGO_CFLAGS}" | grep -- '-g' >/dev/null; then
        CMAKE_DEFS="-DCMAKE_BUILD_TYPE=RelWithDebInfo -DCMAKE_VERBOSE_MAKEFILE=on -DLLAMA_GPROF=on -DLLAMA_SERVER_VERBOSE=on ${CMAKE_DEFS}"
    else
        # TODO - add additional optimization flags...
        CMAKE_DEFS="-DCMAKE_BUILD_TYPE=Release -DLLAMA_SERVER_VERBOSE=off ${CMAKE_DEFS}"
    fi
    case $(uname -s) in
    "Darwin")
        LIB_EXT="dylib"
        WHOLE_ARCHIVE="-Wl,-force_load"
        NO_WHOLE_ARCHIVE=""
        GCC_ARCH="-arch ${ARCH}"
        ;;
    "Linux")
        LIB_EXT="so"
        WHOLE_ARCHIVE="-Wl,--whole-archive"
        NO_WHOLE_ARCHIVE="-Wl,--no-whole-archive"

        # Cross compiling not supported on linux - Use docker
        GCC_ARCH=""
        ;;
    *)
        ;;
    esac
    if [ -z "${CMAKE_CUDA_ARCHITECTURES}" ] ; then
        CMAKE_CUDA_ARCHITECTURES="50;52;61;70;75;80"
    fi
}

# 该函数用于设置 Git 子模块。如果环境变量 OLLAMA_SKIP_PATCHING 被设置，则跳过子模块初始化；
# 否则，先删除旧的子模块目录，然后初始化并强制更新指定的子模块目录。
git_module_setup() {
    if [ -n "${OLLAMA_SKIP_PATCHING}" ]; then
        echo "Skipping submodule initialization"
        return
    fi
    # Make sure the tree is clean after the directory moves
    if [ -d "${LLAMACPP_DIR}/gguf" ]; then
        echo "Cleaning up old submodule"
        rm -rf ${LLAMACPP_DIR}
    fi
    git submodule init
    git submodule update --force ${LLAMACPP_DIR}

}

# 该函数的作用是为 LLAMACPP_DIR 目录下的源代码应用补丁。

# 首先，函数检查 LLAMACPP_DIR 目录中的 CMakeLists.txt 文件是否包含 "ollama" 字符串。如果没有，则向文件末尾追加这行代码：add_subdirectory(../ext_server ext_server) # ollama
# 接着，函数检查 ../patches 目录下是否存在以 .diff 为后缀的补丁文件。如果存在，则依次执行以下操作：

# 对于每个补丁文件，函数首先提取出补丁中要修改的文件列表。
# 然后，对每个文件，函数使用 git checkout 命令将 LLAMACPP_DIR 目录下对应的文件切换到原始状态。
# 最后，函数使用 git apply 命令将补丁应用到 LLAMACPP_DIR 目录下的源代码中。
apply_patches() {
    # Wire up our CMakefile
    if ! grep ollama ${LLAMACPP_DIR}/CMakeLists.txt; then
        echo 'add_subdirectory(../ext_server ext_server) # ollama' >>${LLAMACPP_DIR}/CMakeLists.txt
    fi

    if [ -n "$(ls -A ../patches/*.diff)" ]; then
        # apply temporary patches until fix is upstream
        for patch in ../patches/*.diff; do
            for file in $(grep "^+++ " ${patch} | cut -f2 -d' ' | cut -f2- -d/); do
                (cd ${LLAMACPP_DIR}; git checkout ${file})
            done
        done
        for patch in ../patches/*.diff; do
            (cd ${LLAMACPP_DIR} && git apply ${patch})
        done
    fi
}

# 这个函数名为build()，用于构建项目。它包含两个步骤：

# 使用 cmake 命令在 ${LLAMACPP_DIR} 目录下生成构建文件，输出到 ${BUILD_DIR} 目录中，同时将 ${CMAKE_DEFS} 作为额外的定义参数传递给 cmake。
# 使用 cmake --build 命令在 ${BUILD_DIR} 目录下构建项目，指定构建目标为 ${CMAKE_TARGETS}，并使用 8 个并行进程进行构建。
build() {
    cmake -S ${LLAMACPP_DIR} -B ${BUILD_DIR} ${CMAKE_DEFS}
    cmake --build ${BUILD_DIR} ${CMAKE_TARGETS} -j8
}

# 该函数用于压缩指定目录下的文件以减小二进制文件的整体大小。
# 它首先删除已存在的 gz 文件，然后使用 gzip 命令以最高压缩级别压缩 BUILD_DIR/bin 目录下的所有文件，并将压缩进程的 pid 保存到 pids 变量中。
# 接着，它检查 BUILD_DIR/lib 目录是否存在，如果存在，同样压缩该目录下的所有文件。
# 最后，函数等待所有压缩进程完成，并输出压缩完成的消息。
compress() {
    echo "Compressing payloads to reduce overall binary size..."
    pids=""
    rm -rf ${BUILD_DIR}/bin/*.gz
    for f in ${BUILD_DIR}/bin/* ; do
        gzip -n --best -f ${f} &
        pids+=" $!"
    done
    # check for lib directory
    if [ -d ${BUILD_DIR}/lib ]; then
        for f in ${BUILD_DIR}/lib/* ; do
            gzip -n --best -f ${f} &
            pids+=" $!"
        done
    fi
    echo
    for pid in ${pids}; do
        wait $pid
    done
    echo "Finished compression"
}

# Keep the local tree clean after we're done with the build
# 该函数用于清理 LLAMACPP_DIR 目录下的代码库，将 CMakeLists.txt 文件切换到指定的分支，并根据补丁文件（../patches/*.diff）的内容，将相应的文件切换到代码库中的指定版本。

# 首先，通过 (cd ${LLAMACPP_DIR}/ && git checkout CMakeLists.txt) 命令，切换到 LLAMACPP_DIR 目录，并将 CMakeLists.txt 文件切换到指定的分支。

# 然后，检查 ../patches 目录下是否存在补丁文件（*.diff），如果存在，则进行以下操作：
#   遍历每个补丁文件，使用 grep "^+++ " ${patch} | cut -f2 -d' ' | cut -f2- -d/ 命令提取出补丁文件中修改的文件路径。
#   对于每个提取出的文件路径，切换到 LLAMACPP_DIR 目录，并使用 git checkout ${file} 命令将该文件切换到代码库中的指定版本。
cleanup() {
    (cd ${LLAMACPP_DIR}/ && git checkout CMakeLists.txt)

    if [ -n "$(ls -A ../patches/*.diff)" ]; then
        for patch in ../patches/*.diff; do
            for file in $(grep "^+++ " ${patch} | cut -f2 -d' ' | cut -f2- -d/); do
                (cd ${LLAMACPP_DIR}; git checkout ${file})
            done
        done
    fi
}
