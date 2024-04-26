import imageCompression from "browser-image-compression";

export async function handleImageCopression(file: File) {
  const imageFile = file
  const options = {
    maxSizeMB: 1,
    maxWidthOrHeight: 1920,
    useWebWorker: true,
  }
  try {
    const compressedFile = await imageCompression(imageFile, options);

    return { ok: true, file: compressedFile }
  } catch (error) {
    console.log("ERROR >>>", error);
    if (error instanceof Error) {
      return ({ ok: false, error: error, message: error.message })

    } else {
      return { ok: false, error: error, message: "Misc Error" }
    }
  }

}

export async function uploadImages(formData: FormData) {

  const response = await fetch(window.location.href, { // Or use a specific path if necessary
    method: 'POST',
    body: formData,
    headers: {
      'accept': 'application/json',
    },
  });
  if (!response.ok) {
    console.log("NETWORK RESPONSE WAS WACK")
    throw new Error('Network response was not ok');
  }
}