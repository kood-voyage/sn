import { handleImageCopression, uploadImages } from "$lib/client/image-compression";

export async function handleSubmit(formData: FormData) {

  const image1 = formData.get('image') as File

  if (image1) {
    const compressedImage1 = await handleImageCopression(image1)

    if (compressedImage1.ok && compressedImage1.file) {
      formData.set('image', compressedImage1.file, compressedImage1.file.name);

    } else {
      console.log("nop1")
      formData.set('image', image1, image1.name)
    }
  }

  uploadImages(formData);
}


