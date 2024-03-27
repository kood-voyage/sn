import imageCompression from "browser-image-compression";

export async function handleImageCopression(file: File) {
  console.log("Started handleImage!")
  const imageFile = file
  // console.log('originalFile instanceof Blob', imageFile instanceof Blob); // true
  console.log(`originalFile size ${imageFile.size / 1024 / 1024} MB`);

  const options = {
    maxSizeMB: 1,
    maxWidthOrHeight: 1920,
    useWebWorker: true,
  }
  try {
    const compressedFile = await imageCompression(imageFile, options);
    console.log(`compressedFile size ${compressedFile.size / 1024 / 1024} MB`);

    return { ok: true, file: compressedFile }
  } catch (error) {
    console.log("ERROR >>>", error);
    if (error instanceof Error) {
      return { ok: false, error: error, message: error.message }

    } else {
      return { ok: false, error: error, message: "Misc Error" }
    }
  }

}

export async function uploadImages(formData: FormData) {
  const response = await fetch(window.location.href, {
    method: 'POST',
    body: formData,
    headers: {
      'accept': 'application/json',
    },
  });
  if (!response.ok) {
    throw new Error('Network response was not ok');
  }
}