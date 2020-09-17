build:
	asciidoctor manage-accesskey-vault.adoc -o docs/manage-accesskey-vault.html
	asciidoctor vpc-over-session-manager.adoc -o docs/vpc-over-session-manager.html
	asciidoctor ssh-over-ssm.adoc -o docs/ssh-over-ssm.html
