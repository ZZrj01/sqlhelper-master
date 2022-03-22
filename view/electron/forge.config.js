const FileManagerPlugin = require('filemanager-webpack-plugin');

// 服务端需要用到的文件
let extraResource = [
    "../../serve/bin/data.yml",
    "../../serve/bin/config.yml",
]
let icon; // 图标
// 根据不同平台指定不同文件
if (process.platform === 'darwin') {
    extraResource.push("../../serve/bin/serve")
    icon = "./icon.icns"
} else {
    extraResource.push("../../serve/bin/serve.exe")
    icon = "./icon.ico"
}

// 导出配置
module.exports = {
    packagerConfig: {
        icon: icon,
        overwrite: true, // 是否覆盖
        asar: true, // 是否压缩
        extraResource: extraResource
    },
    makers: [{
            name: "@electron-forge/maker-squirrel",
            config: {
                name: "sqlhelper"
            }
        },
        {
            name: "@electron-forge/maker-zip",
            platforms: [
                "darwin"
            ]
        },
        {
            name: "@electron-forge/maker-deb",
            config: {}
        },
        {
            name: "@electron-forge/maker-rpm",
            config: {}
        }
    ]
}