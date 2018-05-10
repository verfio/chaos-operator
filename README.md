# chaos-operator
## Kubernetes Chaos Operator

Create chaos-operator.yaml file and apply it:

`kubectl apply -f chaos-operator.yaml`


```

apiVersion: apps/v1
kind: Deployment
metadata:
  name: chaos-operator
  labels:
    app: chaos-operator
spec:
  replicas: 1
  selector:
    matchLabels:
      app: chaos-operator
  template:
    metadata:
      labels:
        app: chaos-operator
    spec:
      serviceAccount: chaos-operator
      containers:
      - name: chaos-operator
        image: verfio/chaos-operator
        imagePullPolicy: Always
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: chaos-operator
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRoleBinding
metadata:
  name: chaos-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
  - kind: ServiceAccount
    name: chaos-operator
    namespace: default
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: chaoses.verf.io
spec:
  group: verf.io
  version: v1
  names:
    kind: Chaos
    plural: chaoses
  scope: Namespaced

```


Now you can create Custom  Resources, for example chaos.yaml:


`kubectl apply -f chaos.yaml`

```

apiVersion: verf.io/v1
kind: Chaos
metadata:
  name: example-chaos
spec:
  namespace: chaos

```


Now, every minute one random pod will be deleted in the given namespace (chaos).


Run nginx containers for test:


`kubectl run nginx --image=nginx --replicas=2 --namespace=chaos`


Check chaos-operator logs to follow the chaos activity:


`kubectl logs chaos-operator-78d9fc8696-jl4jk`


```
time="2018-05-09T17:45:10Z" level=info msg="Successfully constructed k8s client"
time="2018-05-09T17:45:10Z" level=info msg="Controller.Run: initiating"
time="2018-05-09T17:45:10Z" level=info msg="Add myresource: default/chaos"
time="2018-05-09T17:45:10Z" level=info msg="Controller.Run: cache sync complete"
time="2018-05-09T17:45:10Z" level=info msg="Controller.runWorker: starting"
time="2018-05-09T17:45:10Z" level=info msg="Controller.processNextItem: start"
time="2018-05-09T17:45:10Z" level=info msg="Controller.processNextItem: object created detected: default/chaos"
time="2018-05-09T17:45:10Z" level=info msg=TestHandler.ObjectCreated
time="2018-05-09T17:45:10Z" level=info msg="new chaos is scheduled" namespace=chaos
time="2018-05-09T17:46:00Z" level=info msg="Pod deleted" namespace=chaos pod=nginx-65899c769f-r88t5
time="2018-05-09T17:47:00Z" level=info msg="Pod deleted" namespace=chaos pod=nginx-65899c769f-l4sx6
```



Check current list of chaos tasks:


`kubectl get chaos`


Delete chaos tasks:


`kubectl delete chaos "name of chaos"`


