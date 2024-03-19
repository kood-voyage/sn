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
    // console.log('compressedFile instanceof Blob', compressedFile instanceof Blob); // true
    console.log(`compressedFile size ${compressedFile.size / 1024 / 1024} MB`); // smaller than maxSizeMB

    return { ok: true, file: compressedFile } // write your own logic
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
  console.log("FORMDATA >>>", formData)
  const response = await fetch(window.location.href, { // Or use a specific path if necessary
    method: 'POST',
    body: formData, // FormData containing the compressed images and other form data
    headers: {
      'accept': 'application/json', // Expecting JSON response
    },
  });
  if (!response.ok) {
    // Handle errors, such as by displaying a message to the user
    throw new Error('Network response was not ok');
  }
  const result = await response.json();
  console.log(result); // Process the response from your SvelteKit action
}