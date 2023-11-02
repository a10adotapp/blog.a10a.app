import { decodeBase64Url } from "./helpers.js";

export async function createCredentials() {
  const credentails = await navigator.credentials.create({
    publicKey: {
      rp: {
        id: "f87f-2001-f74-3a0-6300-dd86-9cad-ddb3-1048.ngrok-free.app",
        name: "a10a localhost",
      },
      user: {
        id: decodeBase64Url("some-user-id"),
        name: "mail@a10a.app",
        displayName: "mail@a10a.app",
      },
      challenge: decodeBase64Url("some-challenge"),
      pubKeyCredParams: [
        {
          type: "public-key",
          alg: -7,
        },
        {
          type: "public-key",
          alg: -257,
        },
      ],
      authenticatorSelection: {
        authenticatorAttachment: "platform",
        residentKey: "required",
        userVerification: "preferred",
      },
      extensions: {
        credProps: true,
      },
    },
  });

  console.log({ credentails });
}
