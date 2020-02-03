# ecs-power

## Setup
Add three SSM parameters for each cluster with the following keys:
```
/<clusterName>/ecs-cluster/AUTOSCALE_MIN
/<clusterName>/ecs-cluster/AUTOSCALE_MAX
/<clusterName>/ecs-cluster/AUTOSCALE_DESIRED

# SSM Parameter Key example for a cluster named "dev":
/dev/ecs-cluster/AUTOSCALE_MIN
```

Configure an autoscaling group in AWS EC2 with a name in the following format
```
<clusterName>-ecs

# for a cluster named "dev" the autoscaling group would be named
dev-ecs
```

## Usage
```
ecs-power --help

# Scale up (turn on) environment to preset values
ecs-power on -c dev

# Shut off ecs-cluster
ecs-power down -c dev
```
