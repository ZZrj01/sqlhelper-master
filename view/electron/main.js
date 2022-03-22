const {
    app,
    BrowserWindow,
    Menu
} = require('electron')
// 路径
const path = require('path')
// 加载窗口
let loadWin;
// 线程得pid 默认-1
let pid = -1
// 执行程序对象
const exec = require('child_process').exec
const spawn = require('child_process').spawn
// 本地需要启动的后台服务名称
let workerProcess
// 主窗体
const createWindow = () => {
    // 隐藏菜单
    Menu.setApplicationMenu(null)
    // 创建主窗口
    const win = new BrowserWindow({
        width: 800,
        height: 600,
        show: false,
        resizable: false,
        allowRunningInsecureContent: true,
        experimentalCanvasFeatures: true,
    })
    // 浏览器显示工具
    // win.webContents.openDevTools();
    // 开发环境只向本地
    if (process.env.NODE_ENV === "dev") {
        win.loadURL("http://127.0.0.1:8080")
    } else {
        win.loadFile("index.html")
    }
    // 首页加载玩就关闭加载页
    // 准备好时显示主窗体
    win.on('ready-to-show', function () {
        win.show() // 初始化后再显示
        loadWin.close() // 关闭加载页
    })
}
// 加载窗
const createLoad = () => {
    const win = new BrowserWindow({
        width: 400,
        height: 300,
        frame: false,
        show: false,
        resizable: false,
        movable: false,
    })
    // 本地加载文件
    win.loadFile("load.html")
    // 加载成功就显示
    win.on('ready-to-show', function () {
        win.show()
    })
    loadWin = win
}
// app 准备好了将要显示
app.whenReady().then(() => {
    // 不是开发环境时候执行本地服务
    if (process.env.NODE_ENV != 'dev') {
        // 执行exe  
        runCmd()
    }
    // 创建窗体
    createLoad()
    // 创建主窗体
    createWindow()
})
// 除了 macOS 外，当所有窗口都被关闭的时候退出程序。 因此，通常对程序和它们在
// 任务栏上的图标来说，应当保持活跃状态，直到用户使用 Cmd + Q 退出。
app.on('window-all-closed', function () {
    console.log("完全退出");
    // workerProcess 存在跟pid != -1
    if (workerProcess && pid != -1) {
        process.kill(workerProcess.pid)
    }
    app.quit()
})

// 执行后台程序
function runCmd() {
    // 后台程序的名字
    const cmdName = "serve"
    const cmdPath = process.platform === 'darwin' ? "./" + cmdName : cmdName + ".exe"
    // 开发环境
    if (process.env.NODE_ENV === "serve") {
        pathStr = path.join(__dirname, "../../serve/bin")
    } else {
        pathStr = process.resourcesPath
    }
    // 执行
    workerProcess = spawn(cmdPath, [], {
        cwd: pathStr
    })
    // 打印正常的后台可执行程序输出
    workerProcess.stdout.on('data', function (data) {
        console.log("pid:" + workerProcess.pid);
        console.log('stdout: ' + data)
        pid = workerProcess.pid
    })
    // 打印错误的后台可执行程序输出
    workerProcess.stderr.on('data', function (data) {
        console.log('stderr: ' + data)
        pid = -1
    })
    // 退出之后的输出
    workerProcess.on('close', function (code) {
        console.log('out code：' + code)
    })
}