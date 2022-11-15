$(function () {
    let socket = null;
    let msgBox = $("#chatbox textarea");

    let messages = $("#messages");

    $("#chatbox").submit(function () {
        if (!msgBox.val()) {
            return false;
        }

        if (!socket) {
            alert("エラー：websocket接続が行われていません");
            return false;
        }
        socket.send(msgBox.cal());
        msgBox.val("");
        return false;
    });

    if (!window["WebSocket"]) {
        alert("エラー：websocketに対応していないブラウザです。");
    } else {
        socket = new WebSocket("ws://localhost:8080/room");
        socket.onclose = function () {
            alert("接続が終了しました。");
        }

        socket.onmessage = function (e) {
            messages.append($("<li>").text(e.data));
        }
    }
});