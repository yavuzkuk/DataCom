#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>

#define SERVER_IP "127.0.0.1"
#define SERVER_PORT 8080

int main() {
    int sock;
    struct sockaddr_in serverAddr;
    char buffer[1024];
    char answer[100];

    // Socket oluştur
    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        perror("Socket creation failed");
        exit(1);
    }

    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(SERVER_PORT);
    serverAddr.sin_addr.s_addr = inet_addr(SERVER_IP);

    // Sunucuya bağlan
    if (connect(sock, (struct sockaddr *)&serverAddr, sizeof(serverAddr)) < 0) {
        perror("Connection failed");
        close(sock);
        exit(1);
    }

    // Soruları sırayla al ve cevapla
    while (1) {
        bzero(buffer, sizeof(buffer));
        int n = read(sock, buffer, sizeof(buffer));
        if (n <= 0) {
            break;
        }
        printf("%s", buffer);  // Sunucudan gelen soruyu ekrana yaz

        printf("Your answer: ");
        fgets(answer, sizeof(answer), stdin);
        write(sock, answer, strlen(answer));  // Cevabını sunucuya gönder

        // Sunucudan doğru/yanlış bildirimini al
        bzero(buffer, sizeof(buffer));
        read(sock, buffer, sizeof(buffer));
        printf("%s", buffer);  // Doğru/yanlış bildirimini yazdır
    }

    // Sonuçları al
    bzero(buffer, sizeof(buffer));
    read(sock, buffer, sizeof(buffer));
    printf("%s", buffer);  // Sonuçları yazdır

    close(sock);
    return 0;
}
