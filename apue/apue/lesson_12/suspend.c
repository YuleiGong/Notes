#include "apue.h"
#include <pthread.h>

int quitflag;
sigset_t mask;

pthread_mutex_t lock = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t waitloc = PTHREAD_COND_INITIALIZER;

void *
thr_fn(void *arg)
{
    int err, signo;

    for (;;){
        err = sigwait(&mask, &signo);
        if (err != 0)
            err_exit(err, "sigwait failed");
        switch (signo){
            case SIGINT:
                printf("\ninterupt\n");
                break;
            case SIGQUIT:
                pthread_mutex_lock(&lock);
                quitflag = 1;
                pthread_mutex_unlock(&lock);
                pthread_cond_signal(&waitloc);
                return (0);
            default:
                printf("unexpected signal %d\n",signo);
                exit(1);
        }
    }

}


int 
main(void)
{
    int err;
    sigset_t oldmask;
    pthread_t tid;

    sigemptyset(&mask); //初始化信号信号集置空
    sigaddset(&mask,SIGINT);//增加SIGINT型号至信号集 程序终止Ctrl-C)
    sigaddset(&mask,SIGQUIT);//和SIGINT类似, 但由QUIT字符(通常是Ctrl-\)来控制
    //添加信号集到信号屏蔽字
    if ((err = pthread_sigmask(SIG_BLOCK,&mask,&oldmask)) != 0)
        err_exit(err, "SIG_BLOCK error");

    err = pthread_create(&tid, NULL, thr_fn, 0);
    if (err != 0)
        err_exit(err, "can't create thread");

    pthread_mutex_lock(&lock); //获取互斥锁
    while (quitflag == 0)
        pthread_cond_wait(&waitloc, &lock);
    pthread_mutex_unlock(&lock); 

    quitflag = 0;

    if (sigprocmask(SIG_SETMASK, &oldmask,NULL) < 1)
        err_sys("SIG_SETMASK error");
    exit(0);

}
