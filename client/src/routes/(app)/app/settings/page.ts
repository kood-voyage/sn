import { handleImageCopression, uploadImages } from "$lib/client/image-compression";

export async function handleSubmit() {
  const form = document.getElementById('imageForm') as HTMLFormElement
  const formData = new FormData(form);

  // Compress and replace each image separately
  const image1 = formData.get('fileInputAvatar') as File
  const image2 = formData.get('fileInputCover') as File
  console.log(image1)
  console.log(image2)

  if (image1) {
    const compressedImage1 = await handleImageCopression(image1)

    if (compressedImage1.ok && compressedImage1.file) {
      formData.set('fileInputAvatar', compressedImage1.file, compressedImage1.file.name);

    } else {
      console.log("nop1")
      formData.set('fileInputAvatar', image1, image1.name)
    }
  }

  if (image2) {
    const compressedImage2 = await handleImageCopression(image2)

    if (compressedImage2.ok && compressedImage2.file) {
      formData.set('fileInputCover', compressedImage2.file, compressedImage2.file.name);

    } else {
      console.log("nop2")
      formData.set('fileInputCover', image2, image2.name)
    }

  }

  // Proceed to upload images
  uploadImages(formData);
}


