import websocket
import time
import threading

# 连接 URL 替换为你的 WebSocket 服务器地址
ws_url = "ws://127.0.0.1:9101/connect"  # 替换为你的服务器地址

# 替换为你的 Authorization Token
auth_token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6ImYwZTQ4NzZlLWE1YjYtMTFlZi1iMGEwLTRhMjQ0NzNhNGU0YyIsImV4cCI6MTczMzE0ODY2MX0.gxl8phpVfRWXycAen_Q6v2bqNW1ZSmzVos2Z8DWCl1E"

def on_message(ws, message):
    print(f"收到消息: {message}")

def on_error(ws, error):
    print(f"连接错误: {error}")

def on_close(ws, close_status_code, close_msg):
    print(f"连接关闭，状态码: {close_status_code}, 消息: {close_msg}")

def on_open(ws):
    print("连接成功！")
    def keep_connection():
        try:
            while True:
                ws.send("ping")  # 模拟心跳包
                time.sleep(30)  # 每30秒发送一次
        except Exception as e:
            print(f"心跳发送出错: {e}")
    threading.Thread(target=keep_connection).start()

if __name__ == "__main__":
    websocket.enableTrace(True)  # 启用调试日志
    headers = {
        "Authorization": auth_token  # 添加 Authorization 头
    }
    ws = websocket.WebSocketApp(
        ws_url,
        header=headers,  # 添加自定义头部
        on_open=on_open,
        on_message=on_message,
        on_error=on_error,
        on_close=on_close,
    )
    try:
        ws.run_forever()  # 连接到服务器并保持运行
    except KeyboardInterrupt:
        print("手动中断")
