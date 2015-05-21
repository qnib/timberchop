# timbershop

Framework to test and (someday use) GROK patterns. It should be seen as a work-in-progress, since I am not yet sure 
how this project is best structured. :)

For now I am using it to verify collected grok patterns and build packages out of them.
Starting with:

- [SLURM patterns](https://github.com/qnib/grok-slurm)
- [elasticsearch patterns](https://github.com/qnib/grok-elasticsearch)

## Setup

```
$ git clone https://github.com/qnib/timberchop
$ cd timberchop
$ go get -d
$ go get gopkg.in/yaml.v2
```

## Test new GROK

If you are developing your grok-patterns might not be in `/etc/qnib/grok/`. You have to define the path...

```
export GROK_BASE=${HOME}/git/grok-slurm/etc/qnib/grok/
```

Afterwards just run the go test.

```
> Using YAML test slurm.yml... task_tuple, ipv4, client_tuple, task_launch_complete, run_prolog, start_batch [OK]
PASS
ok      _/Users/kniepbert/git/timberchop    0.016s
```

It checks out `${GROK_BASE}/tests/*.yml` to find tests to run and executes them. A testfile looks like this:

```
$ cat ${GROK_BASE}/tests/slurm.yml
description: patterns to deal with SLURM logs
owner: "Christian Kniep <christian@qnib.org>"
tests:
  task_tuple:
     compare: "%{SLURM_TASK_TUPLE}"
     input: "5.2"
     result: {
        "slurm_jobid": 5, "slurm_taskid": 2
     }
  ipv4:
     compare: "%{IPV4:ip}"
     input: "172.17.0.12"
     result: {
        "ip": "172.17.0.12"
     }
  client_tuple:
     compare: "%{SLURM_CLIENT_TUPLE}"
     input: "3005.3005@172.17.0.12"
     result: {
        "userid": 3005, "groupid": 3005,
        "slurm_client": "172.17.0.12"
     }
  task_launch_complete:
     compare: "%{SLURM_TASK}"
     input: "slurmd slurmd: launch task 5.2 request from 3005.3005@172.17.0.12 (port 22717)"
     result: {
        "slurm_jobid": 5, "slurm_taskid": 2,
        "userid": 3005, "groupid": 3005,
        "slurmd_port": 22717, "slurm_client": "172.17.0.12"
     }
  run_prolog:
     compare: "%{SLURM_RUN_PROLOG}"
     input: "slurmd slurmd: _run_prolog: run job script took usec=40"
     result: {
        "prolog_wall_usec": 40
     }
  start_batch:
     compare: "%{SLURM_START_BATCH}"
     input: "slurmd slurmd: Launching batch job 4 for UID 4002"
     result: {
        "userid": 4002, "slurm_jobid": 4
     }
```
