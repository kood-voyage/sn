import imageCompression from "browser-image-compression";


export async function handleImageCopression(file: File) {
  const options = {
    maxSizeMB: 0.5,
    maxWidthOrHeight: 1920,
    useWebWorker: true,
  }
  try {
    const compressedFile = await imageCompression(file, options);


    // Create a new File object with compressed content and desired name
    const compressedFileWithName = new File([compressedFile], file.name, {
      type: compressedFile.type,
    });

    return { ok: true, file: compressedFileWithName };
  } catch (error) {
    if (error instanceof Error) {
      return ({ ok: false, error: error, message: error.message })

    } else {
      return { ok: false, error: error, message: "Misc Error" }
    }
  }

}

// export async function uploadImages(formData: FormData) {

//   const response = await fetch(window.location.href, { // Or use a specific path if necessary
//     method: 'POST',
//     body: formData,
//     headers: {
//       'accept': 'application/json',
//     },
//   });
//   if (!response.ok) {
//     console.log("NETWORK RESPONSE WAS WACK")
//     throw new Error('Network response was not ok');
//   }
// }