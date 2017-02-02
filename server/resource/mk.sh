#!/bin/bash

cat << EOF > html.go
package resource

const(
        TemplateText=\`
EOF

cat top.html >> html.go

cat <<EOF >> html.go
\`
      JSText=\`
EOF

cat util.js >> html.go

cat <<EOF >> html.go
\`
)
EOF
