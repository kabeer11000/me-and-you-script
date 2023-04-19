import chalk from "chalk";
import fs from "fs";
import path from "path"
import * as http from "https";
// import {fileURLToPath} from 'url';
import ProgressBar from "progress";
import {setWallpaper} from "wallpaper";
import open from "open";

// const __filename = fileURLToPath(import.meta.url);
const dir = __dirname //|| dirname(fileURLToPath(import.meta.url));

const log = console.log;
(async function () {
    // const ProgressBar =  (await import("progress"));
    // const {setWallpaper} = (await import("wallpaper"));
    // const open = (await import("open"));
    // const chalk = (await import("chalk"));


    function downloadWallpaperFile() {
        const file = fs.createWriteStream(path.join(dir, ".caches/me-and-laiba.png"));
        return new Promise((resolve, reject) => {
            const request = http.get("https://docs.cloud.kabeers.network/c/v/643fd05bd08eb---Group%201.png", function (response) {
                const len = parseInt(response.headers['content-length'], 10);
                const bar = new ProgressBar(':bar', {
                    complete: '=',
                    incomplete: ' ',
                    width: 20,
                    total: len
                });
                response.on("data", (chunk) =>
                    bar.tick(chunk.length))

                response.pipe(file);
                // after download completed close filestream
                file.on("finish", () => {
                    file.close();
                    request.end();
                    resolve(path.join(dir, ".caches/me-and-laiba.png"));
                });
            });
        });
    }

    log(chalk.bold.greenBright('Hi Laiba, I Love You 🥰\nThis is a little program I wrote for you'));
    const cacheDir = path.join(dir, ".caches");
    if (!fs.existsSync(cacheDir)) {
        log(chalk.dim("Downloading our custom wallpaper"));
        await fs.promises.mkdir(path.join(dir, ".caches"));
    }
    if (!fs.existsSync(path.join(dir, ".caches/me-and-laiba.png")))
        /** Create wallpaper cache **/
        await downloadWallpaperFile();

    await setWallpaper(path.join(dir, ".caches/me-and-laiba.png"));
    log(chalk.dim.green("Go check your wallpaper!"));
    setTimeout(() => open("https://i.pinimg.com/564x/d4/14/88/d4148825522fcdf4c99298d5954c5ea2.jpg").then(() => log(chalk.red.dim("that image thing was also me"))), 4 * 60 * 100);
})();

