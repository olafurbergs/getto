# Get To(ken)
A command line tool to get a token from a oauth2 server through authorization flow (browser flow)

## Install

```
wget -O ~/bin/getto https://github.com/olafurbergs/getto/releases/download/{version}/getto_linux_amd64 && chmod u+x ~/bin/getto
```

## Usage

```
GetTo(ken) utility to get access tokens through oauth2 authorization flow
Usage of ./getto:
  -config string
    	The configuration file to use. (default ".getto")
  -init
    	Used to create a configuration file with example configuration profile.
  -print-config
    	Prints out current configuration file.
  -profile string
    	The configuration profile to use. (default "default")
```

## Configuration file example

```
profiles:
  google:
    pkce: false
    authorizationUrl: https://accounts.google.com/o/oauth2/v2/auth
    clientId: 589478345472-afh60s6u44ikncmonmtk4h6rl23ipdr8.apps.googleusercontent.com
    clientSecret: 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
    scopes: https%3A//www.googleapis.com/auth/drive.metadata.readonly
    tokenUrl: https://oauth2.googleapis.com/token
    params:
      prompt: consent
      foo: bar
  dodo:
    pkce: true
    authorizationUrl: https://auth.dodo.com/oauth2/v2/auth
    clientId: myFantasticId
    scopes: openid read profile
    tokenUrl: https://oauth2.dodo.com/token
    params:
      foo: bar
```

## Example usage 

```
~/ getto -profile dodo
{
  id_token eyJraWQiOiItMTQ5NzQxMDcwMSIsIng1dCI6IjU5ai1pSVBzY2ZXUm5ITVFHSEdnb0pEQXIwUSIsImFsZyI6IlJTMjU2In0.eyJleHAiOjE3MDA3Nzk2MDAsIm5iZiI6MTcwMDc3NjAwMCwianRpIjoiMjdmOTJlODYtNTZiNC00Y2Y4LTliZTMtNjY4YjU4ZWQwZTU4IiwiaXNzIjoiaHR0cHM6Ly9sb2dpbi1kZW1vLmN1cml0eS5pby9vYXV0aC92Mi9vYXV0aC1hbm9ueW1vdXMiLCJhdWQiOiJkZW1vLXdlYi1jbGllbnQiLCJzdWIiOiJ6MHJnbHViYiIsImF1dGhfdGltZSI6MTcwMDc3NTk5NiwiaWF0IjoxNzAwNzc2MDAwLCJwdXJwb3NlIjoiaWQiLCJhdF9oYXNoIjoiT2xsaklWbkJUQUZ0ekl2RjFlQnYzQSIsInpvbmVpbmZvIjoiRXVyb3BlL1N0b2NraG9sbSIsIndlYnNpdGUiOiJodHRwczovL2V4YW1wbGUuY29tLyIsImJpcnRoZGF0ZSI6IjE5ODctMDItMjciLCJnZW5kZXIiOiJtYWxlIiwiYW1yIjoidXJuOnNlOmN1cml0eTphdXRoZW50aWNhdGlvbjp1c2VybmFtZTp1c2VybmFtZSIsInByb2ZpbGUiOiJodHRwczovL2V4YW1wbGUuY29tL3owcmdsdWJiIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiejByZ2x1YmIiLCJsb2NhbGUiOiJzdi1TRSIsImdpdmVuX25hbWUiOiJKb2huIiwibWlkZGxlX25hbWUiOiJNaWNoYWVsIiwibm9uY2UiOiIxNTk5MDQ2MTAyNjQ3LWR2NCIsInBpY3R1cmUiOiJodHRwczovL2V4YW1wbGUuY29tL3owcmdsdWJiL3Byb2ZpbGUucG5nIiwic2lkIjoiSWR4TzZRRWY2MVR2MlBFVCIsImFjciI6InVybjpzZTpjdXJpdHk6YXV0aGVudGljYXRpb246dXNlcm5hbWU6dXNlcm5hbWUiLCJkZWxlZ2F0aW9uX2lkIjoiNmFjNGIzNjktMDU1Ni00Yzg4LWI3MTctZjYyYTY1M2RjZGY5Iiwic19oYXNoIjoiTXFUUTJpbEhKdUc1WU4yUUt4NTE2dyIsInVwZGF0ZWRfYXQiOjE0ODkzOTU2MDAsImF6cCI6ImRlbW8td2ViLWNsaWVudCIsIm5pY2tuYW1lIjoiejByZ2x1YmIiLCJuYW1lIjoiSm9obiBNaWNoYWVsIERvZSIsImZhbWlseV9uYW1lIjoiRG9lIn0.JCunqi1cazRRLxpa4F0sZnrJH6eAlg__NXofrzxoEJXZsScJR3WkJZL9HerGWXM2waO0bvRGeiV2F7PStXd4krqFM5lnuqle7boMmqDBTTQ9wFZDQfUaY5IB82eeSxP852PJ2Iw3PiDVEUnj55uK8HeUipgRnSTLmHSyHABvpk3RfZF7d9T672FUqcC-D6ZlH8-EYlFfBpaiEpPSVZjllgtW54YXzkPW_0_3s8skIn1SVBxWdaqdCTkSnlHsyg4dOSSzesBkn8MtV4Udo_CRf0FmvZm7mI08-G1ezbsKENXgDAKPMEP9N7QkxsWaCZfyscmukaa7q5aJCjtiwXXM3Eq9RVUbEUy_vCcVS8KE6-m8hyz0oW5OQvI6cYcmP1PV7wAiJvfj1A1r4rlyLA-YcH11F_ws7xs8kF-y0S8LLW-d3rIZtcyxD2IDYRItFHYzScSTDI6-hGITbxR3ncvmyijePzArexLhEbZYgnLRWgMQpkUT16vQJAGEJEk0WXDMFszHJ1SeG0UEe5KQV8OsUy1ML4bqmaGGxlvGYE_geRRDKANwfgKoqAoSIvzjJZyk5TU0MTSii4WFLhsla3B-j6i673VYKw2SRSlnEVU7xN17QHn-x-Iz0E5tizvIeO1mgvdkzE00iooNP0ZAu9wljEtGNLDwzAUMb1TBDlnChYQ
  token_type bearer
  access_token _0XBPWQQ_3cc99992-cfc8-478c-9a91-c6975032a108
  refresh_token _1XBPWQQ_c1fbd9f4-f85a-4897-aada-4493988a4021
  scope openid read profile
  expires_in 299
}

_0XBPWQQ_3cc99992-cfc8-478c-9a91-c6975032a108
```
## License

MIT