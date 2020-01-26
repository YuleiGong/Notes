#include "apue.h"
#include <errno.h>

static void
sig_hup(int signo)
{
    printf("SIGHUP received, pid = %ld\n", (long)getpid());
}

static void
pr_ids(char *name)
{
    printf("%s: pid = %ld,ppid = %ld pgrp = %ld, tpgrp = %ld\n", name, (long)getpid(), (long)getppid(), (long)getpgrp(), (long)tcgetpgrp(STDIN_FILENO));
    fflush(stdout);
}

int 
main(void)
{
    char c;
    pid_t pid;

    pr_ids("parent");

    if ((pid = fork()) < 0){
        err_sys("fork error");
    }else if(pid > 0){
        //父进程睡眠 子进程在父进程前运行
        sleep(5);
    }else{
        pr_ids("child");
        //挂断信号处理函数 不是发送挂断信号
        signal(SIGHUP, sig_hup);
        //停止子进程 不是终止 先挂起子进程，让父进程结束，成为孤儿进程
        kill(getpid(), SIGTSTP);
        pr_ids("child");
        if (read(STDIN_FILENO, &c, 1) != 1)
            printf("read error %d on contrnlling TTY\n", errno);
    }
    exit(0);
}


