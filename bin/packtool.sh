#!/usr/bin/env bash
source package.sh
source $HOME/.bash_aliases

build_pkg() {
    echo "* 删除老的编译文件：build/$PROJECT"
    rm -rf build/$PROJECT/*
    #mkdir -p build/$PROJECT
    mkdir -p build/$PROJECT/etc
    mkdir -p build/$PROJECT/bin
    PACKTOOL_DIR=$(dirname $0)
    NONO_DIR=$(cd $PACKTOOL_DIR/..; pwd)

    # 生成版本号
    version=$VERSION.$(date +%Y%m%d%H%M)-$(git branch -v | grep "^*" | awk '{print $2"-"$3} ' | tr / .)
    if [ $(git  status | grep '^\s\+modified:' | wc -l) -ne 0 ]; then
        version=$version-M
    fi
    if [ -n "$JENKINS_URL" ]; then
        build_branch_tag=$(echo $BUILD_GIT_BRANCH_TAG | tr / .)
        version=$VERSION-$(date +%Y%m%d%H%M)-$build_branch_tag-$(git branch -v | grep "^*" | awk '{print $5} ' | tr -d ")" )
    fi

    echo "* 新包版本号: $version"
    echo "${PROJECT}\`${version}" > version

    # 修改记录
    echo "* 生成修改记录"
    git log -10 > build/$PROJECT/changes

    build
    if [ $? -ne 0 ]; then
        echo "* 编译出错"
        exit 1
    fi
    echo "* 编译成功"

    echo "* 打包"
    PKG_NAME="$PROJECT.$version.tar.bz2"
    (
        cp version build/$PROJECT
        if [ -f "Dockerfile" ]; then
          cp Dockerfile build/$PROJECT
        fi
        if [ -f "docker-compose.yml" ]; then
          cp Dockerfile build/$PROJECT
        fi
        cp -r etc/* build/$PROJECT/etc/
        if [ -d src/config ]; then
            cp -r src/config/* etc/
            cp -r src/config/* build/$PROJECT/etc/
        fi
        if [ -d bin ]; then
            cp -r bin/* build/$PROJECT/bin/
        fi
        if [ -d nonoapi ]; then
            cp nonoapi/*.proto build/$PROJECT/
        fi
        cp $PACKTOOL_DIR/color_log.sh build/$PROJECT/bin/
        cd build
        tar -cjf $PKG_NAME $PROJECT
        if [ $? -ne 0 ]; then
            echo "* 打包异常"
            exit 1
        fi
        ls -lh $PKG_NAME
    )

    echo
    echo "DONE: build/$PKG_NAME"
    if [  -d "${HOME}/go-server" ]; then
        if [ "$IS_UNPACK" = true ]; then
            tar -xvf build/$PKG_NAME -C $HOME/go-server
        fi
    fi
}
