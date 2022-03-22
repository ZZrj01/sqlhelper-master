// 文件管理插件
// 打包vue过程中进行拷贝，删除等操作
const FileManagerPlugin = require('filemanager-webpack-plugin');

module.exports = {
    publicPath: './',
    configureWebpack: {
        plugins: [
            new FileManagerPlugin({
                events: {
                    // 打包完成以后
                    onEnd: {
                        copy: [{
                            source: __dirname + '/dist',
                            destination: __dirname + '/electron'
                        }],
                    },
                },
            })
        ],
    }
}