import { ListBucketsCommand, PutObjectCommand, S3Client } from '@aws-sdk/client-s3';
import { AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, PROFILE_MEDIA_BUCKET } from '$env/static/private';

import { getUserIdFromCookie } from '../jwt-handle';
import type { RequestEvent } from '@sveltejs/kit';


const client = new S3Client({
  region: "us-east-1",
  credentials: {
    accessKeyId: AWS_ACCESS_KEY_ID,
    secretAccessKey: AWS_SECRET_ACCESS_KEY
  }
});


const params = {

}

const command = new ListBucketsCommand(params)

// const getKey = (path: string) => `profile/${path}.json`;

export async function saveUserAvatarToS3(event: RequestEvent, file: File) {
  const userResp = getUserIdFromCookie(event)
  if (!userResp.ok) {
    return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  }

  const user_id = userResp.user_id as string

  saveToS3("avatar", user_id, file)
}

export async function saveUserBannerToS3(event: RequestEvent, file: File) {
  // TODO add some sort of abuse prevention
  const userResp = getUserIdFromCookie(event)
  if (!userResp.ok) {
    return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  }

  const user_id = userResp.user_id as string

  saveToS3("banner", user_id, file)

};

async function saveToS3(type: string, userId: string, file: File): Promise<void> {
  const extension = file.name.slice(file.name.lastIndexOf('.'));
  const key = `profile/${userId}/${type}${extension}`;

  // Convert File to ArrayBuffer then to Buffer
  const arrayBuffer = await file.arrayBuffer();
  const body = Buffer.from(arrayBuffer);

  const uploadCommand = new PutObjectCommand({
    Bucket: PROFILE_MEDIA_BUCKET, // Replace with your bucket name
    Key: key,
    Body: body,
  });

  try {
    const response = await client.send(uploadCommand);
    console.log("S3 upload success", response);
  } catch (error) {
    console.error("S3 upload error", error);
  }
}





export async function mainUpload() {
  try {
    const results = await client.send(command);
    console.log(results)
    // results.forEach(function (item, index) {
    //   console.log(item);
    // });
    // process data.
  } catch (error) {
    console.error(error)
    // error handling.
  } finally {
    // finally.
  }
  // console.log(client)
  // console.log(command)
}