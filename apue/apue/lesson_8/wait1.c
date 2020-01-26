#include "apue.h"
#include <sys/wait.h>

int
main(void)
{
    pid_t pid;
    int status;

    //正常退出
    if ((pid = fork()) < 0)
        err_sys("fork error");
    else if (pid == 0)
        exit(7);
    
    if (wait(&status) != pid)
        err_sys("wait error");
    pr_exit(status);

    //异常终止
    if ((pid = fork()) < 0)
        err_sys("fork error");
    else if (pid == 0)
        abort();
    if (wait(&status) != pid)
        err_sys("wait error");
    pr_exit(status);

    exit(0);
}
