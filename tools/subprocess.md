# subprocess
* 使用Python执行shell,有如下两种场景
    * __等待__,直到命令执行完成,一次性获取到命令执行的结果。
    * 命令执行同时,实时的获取命令结果
* 执行一个shell会自动打开三个标准文件
    * stdin
    * stdout
    * stderr
## ls -a 一次获取命令执行结果
* shell=False python执行命令时,调用unix底层Api,```cmd = ['ls', '-a']```列表形式
* shell=True bash执行命令,```cmd = ls -a```
* communicate()将__stdin__发送给执行命令的进程,命令执行完后从__stdout__，和 __stderr__中获取数据,返回一个__tuple(stdout, stderr)__,使用communicate()执行命令会变成__同步__,如果命令一直没有完成或异常结束,主程序会一直阻塞。
* 使用communicate() 需要将stderr 和 stdout 定向到__PIPE__
* 使用__communicate__ 会使主程序阻塞,直到子进程结束,由于communicate() 使用了PIPE处理stdout 和 stderr 。相比于__wait()__,可以避免__死锁__

```python
# __等待__,直到命令执行完成,一次性获取到命令执行的结果。阻塞式
import subprocess

def run_cmd(cmd, shell):
    return subprocess.Popen(
        cmd,
        shell=shell,        
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE
    ).communicate()
if __name__ == '__main__':
    print (run_cmd(['ls', '-a'], False))
    print (run_cmd('ls -a', True))
out:
(py36) ➜  subprocess python ls_a.py ls_a.py
(b'.\n..\nls_a.py\n', b'')
(b'.\n..\nls_a.py\n', b'')
```
## ping  命令执行同时,实时的获取命令结果
* p.stdout 为一个打开的文件描述符,使用迭代器输出,实时获取命令结果
* Popen对象创建后，主程序不会自动等待子进程完成,(__命令实际已经执行结束__)

```python
# ping  命令执行同时,实时的获取命令结果 非阻塞式
import subprocess
def run_cmd(cmd):
    return subprocess.Popen(
        cmd,
        shell=False,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE
    )
if __name__ == '__main__':
    p = run_cmd(['ping', 'zhihu.com'])
    #Popen对象创建后，主程序不会自动等待子进程完成
    for i in iter(p.stdout.readline, ''):
        print (i.strip().decode('utf8'))
out:
(py36) ➜  subprocess python ping.py
PING zhihu.com (118.178.213.186): 56 data bytes
64 bytes from 118.178.213.186: icmp_seq=0 ttl=42 time=40.673 ms
64 bytes from 118.178.213.186: icmp_seq=1 ttl=42 time=37.222 ms
64 bytes from 118.178.213.186: icmp_seq=2 ttl=42 time=30.652 ms
64 bytes from 118.178.213.186: icmp_seq=3 ttl=42 time=29.491 ms
```
## call
* call 会一直等待命令执行完成,成功返回0 失败返回 1

```python
>>> import subprocess
>>> subprocess.call(['ls', '-a'])
0
```

## check_output()
* 一直等待完成,检查__返回码__,如果returncode不为0，抛出subprocess.CalledProcessError，该对象包含有returncode属性和output属性，output属性为标准输出的输出结果，可用try...except...来检查。
* 如果正常退出,__check_output()__会返回标准输出

```python
import subprocess
try:
    res = subprocess.check_output(
        ['cd', '-a'],
        stderr=subprocess.STDOUT,
    )
    print (res.decode('utf8'))
except subprocess.CalledProcessError as e:
    out_bytes = e.output
    code = e.returncode
    print (code)
    print (out_bytes.decode('utf8'))
```

