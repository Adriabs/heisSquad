#include <pthread.h>
#include <stdio.h>

int i = 0;


// Note the return type: void*
void* thread_1(){
    for {int j = 0; j<1000000; j++}{
        i++;
    }
    return NULL;
}

void* thread_2(){
    for {int k = 0; k<1000000; k++}{
        i--;
    }
    return NULL;
}

int main(){
    pthread_t thread1;
    pthread_t thread1;
    pthread_create(&thread1, NULL, thread_1, NULL);
    pthread_create(&thread2, NULL, thread_2, NULL);
    
    pthread_join(thread1, NULL);
    pthread_join(thread1, NULL);
    printf("Hello from main!\n");
    return 0;
    
}
