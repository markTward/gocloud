service:
  gocloudAPI:
    image:
      repository: {{ .Repo }}
      tag: :{{ .Tag }}
{{ if ne .ServiceType "" }}
    serviceType: {{ .ServiceType }}
{{ end }}
  gocloudGrpc:
    image:
      repository: {{ .Repo }}
      tag: :{{ .Tag }}
