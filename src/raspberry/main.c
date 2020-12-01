#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <termios.h>
#include <errno.h>
#include <pthread.h> 
#include <sys/ioctl.h>
#include <MQTTClient.h>
#include <wiringPi.h>

#define CLIENTID       "Raspberrypi"

char* MQTT_ADDRESS = "tcp://localhost:1883";
char* MQTT_PUBLISH_TOPIC = "sensor";
char* MQTT_SUBSCRIBE_TOPIC = "led";

MQTTClient client;

void publish(MQTTClient client, char* topic, char* payload);
int on_message(void *context, char *topicName, int topicLen, MQTTClient_message *message);
void setLedStatus(int state);
void *read_serial(void *vargp);

// Turn led ON or OFF
void setLedStatus(int state) {
    digitalWrite(0, !digitalRead(state));
}

// Publish MQTT topic
void publish(MQTTClient client, char* topic, char* payload) {
    MQTTClient_message pubmsg = MQTTClient_message_initializer;

    pubmsg.payload = payload;
    pubmsg.payloadlen = strlen(pubmsg.payload);
    pubmsg.qos = 2;
    pubmsg.retained = 0;
    MQTTClient_deliveryToken token;
    MQTTClient_publishMessage(client, topic, &pubmsg, &token);
    MQTTClient_waitForCompletion(client, token, 1000L);
}

// Subscribe MQTT topic
int on_message(void *context, char *topicName, int topicLen, MQTTClient_message *message) {
    char* payload = message->payload;

    printf("Mensagem recebida! \n\rTopico: %s Mensagem: %s\n", topicName, payload);

    int state = atoi(payload);
    setLedStatus(state);

    publish(client, MQTT_PUBLISH_TOPIC, payload);
    MQTTClient_freeMessage(&message);
    MQTTClient_free(topicName);
    return 1;
}

// Thread read serial
void *read_serial(void *vargp) {
    int sfd = open("/dev/tty0", O_RDWR | O_NOCTTY); 
 	if (sfd == -1) {
 		printf("Error no is : %d\n", errno);
  		printf("Error description is : %s\n", strerror(errno));
        exit(-1);
 	}

 	struct termios options;
 	tcgetattr(sfd, &options);

	cfsetspeed(&options, B9600);
	cfmakeraw(&options);

	options.c_cflag &= ~CSTOPB;
	options.c_cflag |= CLOCAL;
	options.c_cflag |= CREAD;
	options.c_cc[VTIME]=1; 
	options.c_cc[VMIN]=100;

	tcsetattr(sfd, TCSANOW, &options);
	char buf2[100];
    int bytes;

    while(1) {
        usleep(1000);
        ioctl(sfd, FIONREAD, &bytes);
        if(bytes!=0) {
            read(sfd, buf2, 100);
        }
        printf("%s\n\r", buf2);
    }
    close(sfd);
} 

int main(int argc, char *argv[])
{
    // WiringPi setup
    wiringPiSetup();
    pinMode(0, OUTPUT);

    // Setup MQTT
    int rc;
    MQTTClient_connectOptions conn_opts = MQTTClient_connectOptions_initializer;
    MQTTClient_create(&client, MQTT_ADDRESS, CLIENTID, MQTTCLIENT_PERSISTENCE_NONE, NULL);
    MQTTClient_setCallbacks(client, NULL, NULL, on_message, NULL);

    rc = MQTTClient_connect(client, &conn_opts);

    if (rc != MQTTCLIENT_SUCCESS)
    {
        printf("\n\rFalha na conexao ao broker MQTT. Erro: %d\n", rc);
        exit(-1);
    }

    MQTTClient_subscribe(client, MQTT_SUBSCRIBE_TOPIC, 0);

    pthread_t thread_id; 
    pthread_create(&thread_id, NULL, read_serial, NULL); 

    while(1) { }

    pthread_join(thread_id, NULL); 
}