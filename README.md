# go-signer-example

## Run
**RSA:**
```
$ go run rsa/rsa.go
[*] Private key:
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAtZBLnQ6BhVUxDJA0Gavmvpu6cpHajFJi7js2VmSNqNjHgzBn
d/NrcDTZ3x1jbHGzyzcFzwcJ5F0iYGpl5UrRLRoKWRaE4h4cE7MByT6r3jlfzA9x
...
3JtxrCmGWJtwbs3FdOtusnCCEGtedv7tXdgIeJaI1csk4N3qo2sWg4WqXcwtUYoQ
5E/pmYMCgYEAvyaLfJuTc2VhaxR/Ra4zkO2tqwVs5pURUge44L0lr7CJf3IFeYiL
BP2UoZxVAOhKiyjXPPKy2ogT6UyHGeGTZNxSdUrmvSlEPhVMuECMXjD3amyeWg+F
quL0R9UC/6o0uwfYdHiPPpoRA1h0Bx9u9wsjhFpXHacMxg/gN3ek0OM=
-----END RSA PRIVATE KEY-----

[*] Signature:
44d2ef16d97fe18e144189b575f48b36bd48866a3e5fdbf115fa17c5b9a146b67cb98880a13c2109f0af05f20a360f47df40f1b23d5172f1e0d3edf06e2722d6215595078fd396df46d8017f4b68fb5103e9bd0887d7ded98e7cdf65853b413c87b843d78b2b2a45111241387d6a0b1bb0730923ab4ac602a887385557d995d41e5b6f56d0189270a976ca944d1eb9c64789ddf4cf88bb775f70fc807c2db99239738059d40ed103da555984e7691345effff4b062a9b02663a3ccf7119717cf2ef78957ef6331cf74e74fc31954b7186485a1f677f3463361fe586b519ef49c0d87d642a0d4fa430ce97a08a02dc897390ca056bdcd35a1fa5126afab5e9942
[*] Result:
signature is a valid signature of message from the public key.

```

**ECDSA:**
```
$ go run ecdsa/ecdsa.go
[*] Private key:
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIBFZcHSgpBRTkWDd0AdWEEXsUAW/LsP3msGSPqMiRUUloAoGCCqGSM49
AwEHoUQDQgAEK6krIT8I6gSj28blHABchGlkJCouwed4ktvw6ButRjagaERfb4K/
3lI0eVbiFhMx/9D0H8EqdzPdZVDoMWkmWw==
-----END EC PRIVATE KEY-----

[*] Signature:
3046022100b35641da22e7abe22eb003a8af0ff7d6182b782827815540b2f30a1e278de1d4022100b1420e18b2fccc13c20c77153e748e72ef7792173a8c76d9b8467bb5066ad7ec
[*] Result:
true
```

## References 
- https://pkg.go.dev/crypto/rsa@go1.21.5
- https://pkg.go.dev/crypto/ecdsa@go1.21.5