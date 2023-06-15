# hello-api

```sh
curl localhost/translate/hello
```

Port forwarding

```sh
kubectl port-forward service/hello-api 8080
```


# Ingress troubleshoot
```
2.823810      10 event.go:285] Event(v1.ObjectReference{Kind:"Pod", Namespace:"ingress-nginx", Name:"ingress-nginx-controller-784d54644c-xwtdm", UID:"6cb56ec5-a76a-43c7-91fb-3b46d85dc424", APIVersion:"v1", ResourceVersion:"492", FieldPath:""}): type: 'Normal' reason: 'RELOAD' NGINX reload triggered due to a change in configuration
W0615 06:42:01.985651      10 controller.go:1102] Error obtaining Endpoints for Service "default/hello-api": no object matching key "default/hello-api" in local store
I0615 06:42:02.022978      10 admission.go:149] processed ingress via admission controller {testedIngressLength:1 testedIngressTime:0.037s renderingIngressLength:1 renderingIngressTime:0s admissionTime:18.1kBs testedConfigurationSize:0.037}
I0615 06:42:
```
