// go run mksyscall_aix_ppc64.go -aix -tags aix,ppc64 syscall_aix.go syscall_aix_ppc64.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build aix,ppc64

package unix

import (
	"unsafe"
)

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimes(path string, times *[2]Timeval) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callutimes(uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func utimensat(dirfd int, path string, times *[2]Timespec, flag int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callutimensat(dirfd, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(times)), flag)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getcwd(buf []byte) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, e1 := callgetcwd(uintptr(unsafe.Pointer(_p0)), len(buf))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) {
	r0, e1 := callaccept(s, uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getdirent(fd int, buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, e1 := callgetdirent(fd, uintptr(unsafe.Pointer(_p0)), len(buf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func wait4(pid Pid_t, status *_C_int, options int, rusage *Rusage) (wpid Pid_t, err error) {
	r0, e1 := callwait4(int(pid), uintptr(unsafe.Pointer(status)), options, uintptr(unsafe.Pointer(rusage)))
	wpid = Pid_t(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func ioctl(fd int, req uint, arg uintptr) (err error) {
	_, e1 := callioctl(fd, int(req), arg)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func FcntlInt(fd uintptr, cmd int, arg int) (r int, err error) {
	r0, e1 := callfcntl(fd, cmd, uintptr(arg))
	r = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) (err error) {
	_, e1 := callfcntl(fd, cmd, uintptr(unsafe.Pointer(lk)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func fcntl(fd int, cmd int, arg int) (val int, err error) {
	r0, e1 := callfcntl(uintptr(fd), cmd, uintptr(arg))
	val = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Acct(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callacct(uintptr(unsafe.Pointer(_p0)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chdir(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callchdir(uintptr(unsafe.Pointer(_p0)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Chroot(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callchroot(uintptr(unsafe.Pointer(_p0)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Close(fd int) (err error) {
	_, e1 := callclose(fd)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup(oldfd int) (fd int, err error) {
	r0, e1 := calldup(oldfd)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Exit(code int) {
	callexit(code)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callfaccessat(dirfd, uintptr(unsafe.Pointer(_p0)), mode, flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchdir(fd int) (err error) {
	_, e1 := callfchdir(fd)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmod(fd int, mode uint32) (err error) {
	_, e1 := callfchmod(fd, mode)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callfchmodat(dirfd, uintptr(unsafe.Pointer(_p0)), mode, flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callfchownat(dirfd, uintptr(unsafe.Pointer(_p0)), uid, gid, flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fdatasync(fd int) (err error) {
	_, e1 := callfdatasync(fd)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fsync(fd int) (err error) {
	_, e1 := callfsync(fd)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpgid(pid int) (pgid int, err error) {
	r0, e1 := callgetpgid(pid)
	pgid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpgrp() (pid int) {
	r0, _ := callgetpgrp()
	pid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpid() (pid int) {
	r0, _ := callgetpid()
	pid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getppid() (ppid int) {
	r0, _ := callgetppid()
	ppid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getpriority(which int, who int) (prio int, err error) {
	r0, e1 := callgetpriority(which, who)
	prio = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrusage(who int, rusage *Rusage) (err error) {
	_, e1 := callgetrusage(who, uintptr(unsafe.Pointer(rusage)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getsid(pid int) (sid int, err error) {
	r0, e1 := callgetsid(pid)
	sid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Kill(pid int, sig Signal) (err error) {
	_, e1 := callkill(pid, int(sig))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Klogctl(typ int, buf []byte) (n int, err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	r0, e1 := callsyslog(typ, uintptr(unsafe.Pointer(_p0)), len(buf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdir(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callmkdir(dirfd, uintptr(unsafe.Pointer(_p0)), mode)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callmkdirat(dirfd, uintptr(unsafe.Pointer(_p0)), mode)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mkfifo(path string, mode uint32) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callmkfifo(uintptr(unsafe.Pointer(_p0)), mode)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mknod(path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callmknod(uintptr(unsafe.Pointer(_p0)), mode, dev)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mknodat(dirfd int, path string, mode uint32, dev int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callmknodat(dirfd, uintptr(unsafe.Pointer(_p0)), mode, dev)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Nanosleep(time *Timespec, leftover *Timespec) (err error) {
	_, e1 := callnanosleep(uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Open(path string, mode int, perm uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, e1 := callopen64(uintptr(unsafe.Pointer(_p0)), mode, perm)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	r0, e1 := callopenat(dirfd, uintptr(unsafe.Pointer(_p0)), flags, mode)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func read(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, e1 := callread(fd, uintptr(unsafe.Pointer(_p0)), len(p))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Readlink(path string, buf []byte) (n int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	var _p1 *byte
	if len(buf) > 0 {
		_p1 = &buf[0]
	}
	r0, e1 := callreadlink(uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(_p1)), len(buf))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(oldpath)
	if err != nil {
		return
	}
	var _p1 *byte
	_p1, err = BytePtrFromString(newpath)
	if err != nil {
		return
	}
	_, e1 := callrenameat(olddirfd, uintptr(unsafe.Pointer(_p0)), newdirfd, uintptr(unsafe.Pointer(_p1)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setdomainname(p []byte) (err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	_, e1 := callsetdomainname(uintptr(unsafe.Pointer(_p0)), len(p))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Sethostname(p []byte) (err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	_, e1 := callsethostname(uintptr(unsafe.Pointer(_p0)), len(p))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpgid(pid int, pgid int) (err error) {
	_, e1 := callsetpgid(pid, pgid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setsid() (pid int, err error) {
	r0, e1 := callsetsid()
	pid = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Settimeofday(tv *Timeval) (err error) {
	_, e1 := callsettimeofday(uintptr(unsafe.Pointer(tv)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setuid(uid int) (err error) {
	_, e1 := callsetuid(uid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setgid(uid int) (err error) {
	_, e1 := callsetgid(uid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setpriority(which int, who int, prio int) (err error) {
	_, e1 := callsetpriority(which, who, prio)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Statx(dirfd int, path string, flags int, mask int, stat *Statx_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callstatx(dirfd, uintptr(unsafe.Pointer(_p0)), flags, mask, uintptr(unsafe.Pointer(stat)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Sync() {
	callsync()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Times(tms *Tms) (ticks uintptr, err error) {
	r0, e1 := calltimes(uintptr(unsafe.Pointer(tms)))
	ticks = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Umask(mask int) (oldmask int) {
	r0, _ := callumask(mask)
	oldmask = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Uname(buf *Utsname) (err error) {
	_, e1 := calluname(uintptr(unsafe.Pointer(buf)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlink(path string) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callunlink(uintptr(unsafe.Pointer(_p0)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Unlinkat(dirfd int, path string, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callunlinkat(dirfd, uintptr(unsafe.Pointer(_p0)), flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Ustat(dev int, ubuf *Ustat_t) (err error) {
	_, e1 := callustat(dev, uintptr(unsafe.Pointer(ubuf)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func write(fd int, p []byte) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, e1 := callwrite(fd, uintptr(unsafe.Pointer(_p0)), len(p))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func readlen(fd int, p *byte, np int) (n int, err error) {
	r0, e1 := callread(fd, uintptr(unsafe.Pointer(p)), np)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func writelen(fd int, p *byte, np int) (n int, err error) {
	r0, e1 := callwrite(fd, uintptr(unsafe.Pointer(p)), np)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Dup2(oldfd int, newfd int) (err error) {
	_, e1 := calldup2(oldfd, newfd)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fadvise(fd int, offset int64, length int64, advice int) (err error) {
	_, e1 := callposix_fadvise64(fd, offset, length, advice)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fchown(fd int, uid int, gid int) (err error) {
	_, e1 := callfchown(fd, uid, gid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstat(fd int, stat *Stat_t) (err error) {
	_, e1 := callfstat(fd, uintptr(unsafe.Pointer(stat)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callfstatat(dirfd, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)), flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Fstatfs(fd int, buf *Statfs_t) (err error) {
	_, e1 := callfstatfs(fd, uintptr(unsafe.Pointer(buf)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Ftruncate(fd int, length int64) (err error) {
	_, e1 := callftruncate(fd, length)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getegid() (egid int) {
	r0, _ := callgetegid()
	egid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Geteuid() (euid int) {
	r0, _ := callgeteuid()
	euid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getgid() (gid int) {
	r0, _ := callgetgid()
	gid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getuid() (uid int) {
	r0, _ := callgetuid()
	uid = int(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lchown(path string, uid int, gid int) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := calllchown(uintptr(unsafe.Pointer(_p0)), uid, gid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Listen(s int, n int) (err error) {
	_, e1 := calllisten(s, n)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Lstat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := calllstat(uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pause() (err error) {
	_, e1 := callpause()
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pread(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, e1 := callpread64(fd, uintptr(unsafe.Pointer(_p0)), len(p), offset)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, e1 := callpwrite64(fd, uintptr(unsafe.Pointer(_p0)), len(p), offset)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	r0, e1 := callselect(nfd, uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Pselect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timespec, sigmask *Sigset_t) (n int, err error) {
	r0, e1 := callpselect(nfd, uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), uintptr(unsafe.Pointer(sigmask)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setregid(rgid int, egid int) (err error) {
	_, e1 := callsetregid(rgid, egid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setreuid(ruid int, euid int) (err error) {
	_, e1 := callsetreuid(ruid, euid)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Shutdown(fd int, how int) (err error) {
	_, e1 := callshutdown(fd, how)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error) {
	r0, e1 := callsplice(rfd, uintptr(unsafe.Pointer(roff)), wfd, uintptr(unsafe.Pointer(woff)), len, flags)
	n = int64(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Stat(path string, stat *Stat_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callstat(uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(stat)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Statfs(path string, buf *Statfs_t) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callstatfs(uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(buf)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Truncate(path string, length int64) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := calltruncate(uintptr(unsafe.Pointer(_p0)), length)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, e1 := callbind(s, uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
	_, e1 := callconnect(s, uintptr(addr), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getgroups(n int, list *_Gid_t) (nn int, err error) {
	r0, e1 := callgetgroups(n, uintptr(unsafe.Pointer(list)))
	nn = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setgroups(n int, list *_Gid_t) (err error) {
	_, e1 := callsetgroups(n, uintptr(unsafe.Pointer(list)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
	_, e1 := callgetsockopt(s, level, name, uintptr(val), uintptr(unsafe.Pointer(vallen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
	_, e1 := callsetsockopt(s, level, name, uintptr(val), vallen)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socket(domain int, typ int, proto int) (fd int, err error) {
	r0, e1 := callsocket(domain, typ, proto)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func socketpair(domain int, typ int, proto int, fd *[2]int32) (err error) {
	_, e1 := callsocketpair(domain, typ, proto, uintptr(unsafe.Pointer(fd)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, e1 := callgetpeername(fd, uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
	_, e1 := callgetsockname(fd, uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
	var _p0 *byte
	if len(p) > 0 {
		_p0 = &p[0]
	}
	r0, e1 := callrecvfrom(fd, uintptr(unsafe.Pointer(_p0)), len(p), flags, uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
	var _p0 *byte
	if len(buf) > 0 {
		_p0 = &buf[0]
	}
	_, e1 := callsendto(s, uintptr(unsafe.Pointer(_p0)), len(buf), flags, uintptr(to), uintptr(addrlen))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, e1 := callnrecvmsg(s, uintptr(unsafe.Pointer(msg)), flags)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
	r0, e1 := callnsendmsg(s, uintptr(unsafe.Pointer(msg)), flags)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func munmap(addr uintptr, length uintptr) (err error) {
	_, e1 := callmunmap(addr, length)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Madvise(b []byte, advice int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, e1 := callmadvise(uintptr(unsafe.Pointer(_p0)), len(b), advice)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mprotect(b []byte, prot int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, e1 := callmprotect(uintptr(unsafe.Pointer(_p0)), len(b), prot)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, e1 := callmlock(uintptr(unsafe.Pointer(_p0)), len(b))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Mlockall(flags int) (err error) {
	_, e1 := callmlockall(flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Msync(b []byte, flags int) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, e1 := callmsync(uintptr(unsafe.Pointer(_p0)), len(b), flags)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Munlock(b []byte) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	_, e1 := callmunlock(uintptr(unsafe.Pointer(_p0)), len(b))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Munlockall() (err error) {
	_, e1 := callmunlockall()
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func pipe(p *[2]_C_int) (err error) {
	_, e1 := callpipe(uintptr(unsafe.Pointer(p)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func poll(fds *PollFd, nfds int, timeout int) (n int, err error) {
	r0, e1 := callpoll(uintptr(unsafe.Pointer(fds)), nfds, timeout)
	n = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func gettimeofday(tv *Timeval, tzp *Timezone) (err error) {
	_, e1 := callgettimeofday(uintptr(unsafe.Pointer(tv)), uintptr(unsafe.Pointer(tzp)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Time(t *Time_t) (tt Time_t, err error) {
	r0, e1 := calltime(uintptr(unsafe.Pointer(t)))
	tt = Time_t(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Utime(path string, buf *Utimbuf) (err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(path)
	if err != nil {
		return
	}
	_, e1 := callutime(uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(buf)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getsystemcfg(label int) (n uint64) {
	r0, _ := callgetsystemcfg(label)
	n = uint64(r0)
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Getrlimit(resource int, rlim *Rlimit) (err error) {
	_, e1 := callgetrlimit(resource, uintptr(unsafe.Pointer(rlim)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Setrlimit(resource int, rlim *Rlimit) (err error) {
	_, e1 := callsetrlimit(resource, uintptr(unsafe.Pointer(rlim)))
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Seek(fd int, offset int64, whence int) (off int64, err error) {
	r0, e1 := calllseek(fd, offset, whence)
	off = int64(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) {
	r0, e1 := callmmap64(addr, length, prot, flags, fd, offset)
	xaddr = uintptr(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
