// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>


	int i;
	pthread_mutex_t lock;

// Note the return type: void*
void* thread1Func(){
	int j;
	for (j=0;j<1000000;j++){
		pthread_mutex_lock(&lock);
		i++;
		pthread_mutex_unlock(&lock);
	}
    return NULL;
}

void* thread2Func(){
	int j;
	for (j=0;j<1000000;j++){
		pthread_mutex_lock(&lock);
		i--;
		pthread_mutex_unlock(&lock);
	}
    return NULL;
}


int main(){
	i = 0;
	
	if (pthread_mutex_init(&lock, NULL) != 0)
    {
        printf("\n mutex init failed\n");
        return 1;
    }
	
    pthread_t thread1;
    pthread_create(&thread1, NULL, thread1Func, NULL);
	
	pthread_t thread2;
    pthread_create(&thread2, NULL, thread2Func, NULL);
    
    pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);
    printf("%d", i);
	pthread_mutex_destroy(&lock);
    return 0;
    
}