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

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        perror("Socket creation failed");
        exit(1);
    }

    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(SERVER_PORT);
    serverAddr.sin_addr.s_addr = inet_addr(SERVER_IP);

    if (connect(sock, (struct sockaddr *)&serverAddr, sizeof(serverAddr)) < 0) {
        perror("Connection failed");
        close(sock);
        exit(1);
    }

    int counter = 1;
    while (counter) {
        bzero(buffer, sizeof(buffer));
        int n = read(sock, buffer, sizeof(buffer));
        if (n <= 0) {
            break;
        }


        if(strcmp(buffer, "skip") == 0){
            counter = 0;
            break;
        }

        printf("%s", buffer);

        printf("Your answer: ");
        fgets(answer, sizeof(answer), stdin);
        write(sock, answer, strlen(answer));
        
        bzero(buffer, sizeof(buffer));
        read(sock, buffer, sizeof(buffer));
        printf("%s", buffer);


        if (strcmp(buffer, "skip") == 0) {
            counter = 0;
            break;
        }

    }

    bzero(buffer, sizeof(buffer));
    read(sock, buffer, sizeof(buffer));
    printf("%s", buffer);

    bzero(buffer, sizeof(buffer));
    read(sock, buffer, sizeof(buffer));
    printf("%s", buffer);

    bzero(buffer, sizeof(buffer));
    read(sock, buffer, sizeof(buffer));
    printf("%s", buffer);

    close(sock);
    return 0;
}
