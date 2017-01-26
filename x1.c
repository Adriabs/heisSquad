// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread

#include <pthread.h>
#include <stdio.h>


	int i;

// Note the return type: void*
void* thread1Func(){
	int j;
	for (j=0;j<1000000;j++){
		i++;
	}
    return NULL;
}

void* thread2Func(){
	int j;
	for (j=0;j<1000000;j++){
		i--;
	}
    return NULL;
}


int main(){
	i = 0;
	
    pthread_t thread1;
    pthread_create(&thread1, NULL, thread1Func, NULL);
	
	pthread_t thread2;
    pthread_create(&thread2, NULL, thread2Func, NULL);
    
    pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);
    printf("%d", i);
    return 0;
    
}