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