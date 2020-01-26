#include "apue.h"
int
main(void)
{
    if (chdir("/tmp") < 0)
        err_sys("child failed");
    printf("child to /tmp succeeded\n");
    exit(0);
}
