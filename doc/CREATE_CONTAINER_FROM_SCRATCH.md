* ## Building a Container with a Namespace
Namespaces isolate processes within their own execution sandbox so that they run completely isolated from other processes in different namespaces.
There are six namespaces:

* **PID namespace:** The processes within the PID namespace have a different process tree. They have an init process with a PID of 1.
* **Mount namespace:** This namespace controls which mount points a process can see. If a process is within a namespace, it will only see the mounts within that namespace.
* **UTS namespace:** This allows a process to see a different namespace than the actual global namespace.
* **Network namespace:** This namespace gives a different network view within a namespace. Network constructs like ports, iptables, and so on, are scoped within the namespace.
* **IPC namespace:** This namespace confines interprocess communication structures like pipes within a specific namespace.
* **User-namespace:** This namespace allows for a separate user and group view within the namespace.


* ### Building the project 
```
go mod init container
go mod tidy
go build
# execute using root 
./container
```
* Test results
```
root@system76-pc:/home/mohamed/src/containers# ls -li /proc/self/ns/uts
274509 lrwxrwxrwx 1 root root 0 Dec 19 11:17 /proc/self/ns/uts -> 'uts:[4026534187]'
root@system76-pc:/home/mohamed/src/containers# ./container             
root@system76-pc:/home/mohamed/src/containers# ls -lrt /proc/self/ns/uts
lrwxrwxrwx 1 root root 0 Dec 19 11:17 /proc/self/ns/uts -> 'uts:[4026534188]'
root@system76-pc:/home/mohamed/src/containers# exit
```


* Adding more Namespace. Now we see that the ownership belongs to nobody. This is because we also used a user-namespace as a clone flag. The container is now within a new user-namespace. User-namespaces require that we map the user from the namespace to the host. Since we have not done anything yet, we still see nobody as the user.
```
root@system76-pc:/home/mohamed/src/containers# ls -li /proc/self/ns/ 
total 0
332402 lrwxrwxrwx 1 root root 0 Dec 19 11:40 cgroup -> 'cgroup:[4026531835]'
332397 lrwxrwxrwx 1 root root 0 Dec 19 11:40 ipc -> 'ipc:[4026531839]'
332401 lrwxrwxrwx 1 root root 0 Dec 19 11:40 mnt -> 'mnt:[4026531841]'
332395 lrwxrwxrwx 1 root root 0 Dec 19 11:40 net -> 'net:[4026531840]'
332398 lrwxrwxrwx 1 root root 0 Dec 19 11:40 pid -> 'pid:[4026531836]'
332399 lrwxrwxrwx 1 root root 0 Dec 19 11:40 pid_for_children -> 'pid:[4026531836]'
332403 lrwxrwxrwx 1 root root 0 Dec 19 11:40 time -> 'time:[4026531834]'
332404 lrwxrwxrwx 1 root root 0 Dec 19 11:40 time_for_children -> 'time:[4026531834]'
332400 lrwxrwxrwx 1 root root 0 Dec 19 11:40 user -> 'user:[4026531837]'
332396 lrwxrwxrwx 1 root root 0 Dec 19 11:40 uts -> 'uts:[4026531838]'
root@system76-pc:/home/mohamed/src/containers# go build
root@system76-pc:/home/mohamed/src/containers# ./container 
nobody@system76-pc:/home/mohamed/src/containers$ ls -li /proc/self/ns/
total 0
319542 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 cgroup -> 'cgroup:[4026531835]'
319537 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 ipc -> 'ipc:[4026533884]'
319541 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 mnt -> 'mnt:[4026533882]'
319535 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 net -> 'net:[4026533886]'
319538 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 pid -> 'pid:[4026533885]'
319539 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 pid_for_children -> 'pid:[4026533885]'
319543 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 time -> 'time:[4026531834]'
319544 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 time_for_children -> 'time:[4026531834]'
319540 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 user -> 'user:[4026533881]'
319536 lrwxrwxrwx 1 nobody nogroup 0 Dec 19 11:41 uts -> 'uts:[4026533883]'

```
* After Adding UidMappings and GidMappings
```
# ls -li /proc/self/ns/ 
356052 lrwxrwxrwx 1 root root 0 Dec 19 11:56 cgroup -> 'cgroup:[4026531835]'
356047 lrwxrwxrwx 1 root root 0 Dec 19 11:56 ipc -> 'ipc:[4026531839]'
356051 lrwxrwxrwx 1 root root 0 Dec 19 11:56 mnt -> 'mnt:[4026531841]'
356045 lrwxrwxrwx 1 root root 0 Dec 19 11:56 net -> 'net:[4026531840]'
356048 lrwxrwxrwx 1 root root 0 Dec 19 11:56 pid -> 'pid:[4026531836]'
356049 lrwxrwxrwx 1 root root 0 Dec 19 11:56 pid_for_children -> 'pid:[4026531836]'
356053 lrwxrwxrwx 1 root root 0 Dec 19 11:56 time -> 'time:[4026531834]'
356054 lrwxrwxrwx 1 root root 0 Dec 19 11:56 time_for_children -> 'time:[4026531834]'
356050 lrwxrwxrwx 1 root root 0 Dec 19 11:56 user -> 'user:[4026531837]'
356046 lrwxrwxrwx 1 root root 0 Dec 19 11:56 uts -> 'uts:[4026531838]'
# ./container 
# ls -li /proc/self/ns/ 
367724 lrwxrwxrwx 1 root root 0 Dec 19 11:56 cgroup -> 'cgroup:[4026531835]'
367719 lrwxrwxrwx 1 root root 0 Dec 19 11:56 ipc -> 'ipc:[4026534130]'
367723 lrwxrwxrwx 1 root root 0 Dec 19 11:56 mnt -> 'mnt:[4026534128]'
367717 lrwxrwxrwx 1 root root 0 Dec 19 11:56 net -> 'net:[4026534132]'
367720 lrwxrwxrwx 1 root root 0 Dec 19 11:56 pid -> 'pid:[4026534131]'
367721 lrwxrwxrwx 1 root root 0 Dec 19 11:56 pid_for_children -> 'pid:[4026534131]'
367725 lrwxrwxrwx 1 root root 0 Dec 19 11:56 time -> 'time:[4026531834]'
367726 lrwxrwxrwx 1 root root 0 Dec 19 11:56 time_for_children -> 'time:[4026531834]'
367722 lrwxrwxrwx 1 root root 0 Dec 19 11:56 user -> 'user:[4026534127]'
367718 lrwxrwxrwx 1 root root 0 Dec 19 11:56 uts -> 'uts:[4026534129]'

```
