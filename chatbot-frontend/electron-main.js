// In the main process.
const {app, BrowserWindow } = require('electron')

app.on('ready', () => {
  const mainWindow = new BrowserWindow({ width: 800, height: 600, frame: true });
  mainWindow.loadURL('http://localhost:8000');
});

// Exit app when all windows are closed
app.on('window-all-closed',async () => {
    app.quit();
})

// On macOS it's common to leave the app running until the user explicitly quits (Cmd + Q)
app.on('activate', () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    const mainWindow = new BrowserWindow({ width: 800, height: 600, frame: true });
    mainWindow.loadURL('http://localhost:8000');
  }
});