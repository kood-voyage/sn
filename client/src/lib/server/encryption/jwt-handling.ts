//  // Encode returns signed JWT token
// func (a *Algorithm) Encode(payload *Claims) (string, error) {
// 	header := a.NewHeader()

import type { Claims } from "$lib/types/encryption/claim";

// 	jsonHeader, err := json.Marshal(header)
// 	if err != nil {
// 		return "", err
// 	}

// 	b64Header := base64.RawURLEncoding.EncodeToString(jsonHeader)

// 	jsonPayload, err := json.Marshal(payload.claimsMap)
// 	if err != nil {
// 		return "", err
// 	}

// 	b64Payload := base64.RawURLEncoding.EncodeToString(jsonPayload)

// 	unsignedToken := b64Header + "." + b64Payload

// 	signature, err := a.Sign(unsignedToken)

// 	b64Signature := base64.RawURLEncoding.EncodeToString(signature)

// 	token := b64Header + "." + b64Payload + "." + b64Signature

// 	return token, nil
// }
function getHeader() {
  return {
    Typ: "JWT",
    Alg: "HS256",
  }
}

export function encode(payload: Claims) {
  const header = getHeader()

  const jsonHeader = JSON.stringify(header);
  const b64Header = Buffer.from(jsonHeader).toString('base64url');

  const jsonPayload = JSON.stringify(payload.claimsMap)
  const b64Payload = Buffer.from(jsonPayload).toString('base64url');

  const unsignedToken = b64Header + "." + b64Payload



  console.log(unsignedToken)

}

// How does the signing work, I understand that the write makes the