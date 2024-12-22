const SCREEN_W = 600;
const SCREEN_H = 400;

const ROAD_L = 2000; // roadの長さ
const VIEW_L = 300; // 視界の長さ

const PART_L = 100; // 各partの幅（z方向）
const CAMERA_D = 0.8; // roadとの距離がこの値に等しいときscaleが1になる
const ROAD_W = 1000; // roadの幅（x方向）
const JIKI_Y = 1000; // 自機の高さ

const can = document.getElementById("can");
can.width = SCREEN_W;
can.height = SCREEN_H;
const con = can.getContext("2d");

const road = [];
for (let i = 0; i < ROAD_L; i++) {
    let part = {};
    // partの中心の座標
    part.x = 0; // 横
    part.y = 0; // 高さ
    part.z = i * PART_L; // 奥行き
    part.c = 0; // カーブ係数
    part.sx = 0;

    if (i > 100 && i <= 200) part.c = 1;
    if (i > 300 && i <= 400) part.c = -4;

    if (i > 400) part.y = Math.sin((i - 400) / 30) * 1000;

    road.push(part);
}

let jiki_x = 0;
let jiki_y = JIKI_Y;
let jiki_z = 0;

function drawRoad(col, mx, my, mw, px, py, pw) {
    let x1, y1, x2, y2;

    y1 = my;
    y2 = py;
    x1 = mx - mw / 2;
    x2 = px - pw / 2;

    // console.log(x1, y1, x2, y2);

    con.fillStyle = col;
    // con.fillRect(px - pw / 2, py, pw, 1);
    con.beginPath();
    con.moveTo(x1, y1);
    con.lineTo(x1 + mw, y1);
    con.lineTo(x2 + pw, y2);
    con.lineTo(x2, y2);
    con.closePath();
    con.fill();
}

// setInterval(mainLoop, 1000 / 60);

let mx, my, mw; // 1個前のroadの描画情報
mx = 0;
mainLoop();

function mainLoop() {

    con.fillStyle = "#6af";
    con.fillRect(0, 0, SCREEN_W, SCREEN_H);

    con.fillStyle = "#000";
    con.font = "14px sans-serif";
    con.fillText(`${jiki_x} ${jiki_z}`, 14, 14);

    let start = jiki_z / PART_L;
    jiki_y = JIKI_Y + road[start].y;

    // カーブを計算。sxが三角数になり、2次曲線になる
    // let sx = 0;
    // let cx = 0;
    // for (let i = start; i < start + VIEW_L; i++) {
    //     let r = road[i];
    //     cx += r.c;
    //     sx += cx;
    //     r.sx = sx;
    // }

    // 描画（奥から手前に）
    for (let i = start + VIEW_L - 1; i >= start ; i--) {
        let r = road[i];
        let dist = r.z - jiki_z; // プレイヤーとの距離
        let scale = CAMERA_D / dist;

        let px =  (1 + (r.x - jiki_x + r.sx) * scale) * SCREEN_W / 2;
        let py =  (1 - (r.y - jiki_y) * scale) * SCREEN_H / 2;
        let pw = ROAD_W * scale * SCREEN_W;

        px = Math.floor(px);
        py = Math.floor(py);
        pw = Math.floor(pw);

        console.log(px, py);

        if (mx) {
            let col = (i % 3) ? "#aaa" : "#bbb";
            let edg = (i % 3) ? "#bbb" : "#fff";
            let grn = (i % 5) ? "#6f6" : "#8f8";
            drawRoad(grn, SCREEN_W / 2, my, SCREEN_W, SCREEN_W / 2, py, SCREEN_W);
            drawRoad(edg, mx, my, mw * 1.1, px, py, pw * 1.1);
            drawRoad(col, mx, my, mw, px, py, pw);
        }
        mx = px;
        my = py;
        mw = pw;
    }
}

onkeydown = function (e) {
    switch (e.keyCode) {
        case 37: // 左
            jiki_x -= 100;
            break;
        case 38: // 上
            jiki_z += 100;
            break;
        case 39: // 右
            jiki_x += 100;
            break;
        case 40: // 下
            jiki_z -= 100;
            break;
    }

}