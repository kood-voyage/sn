import { ListBucketsCommand, PutObjectCommand, S3Client } from '@aws-sdk/client-s3';
import { AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, PROFILE_MEDIA_BUCKET } from '$env/static/private';




type UserResp = {
  ok: boolean;
  user_id: string;
  error?: undefined;
  message?: undefined;
} | {
  ok: boolean;
  error: unknown;
  message: string;
  user_id?: undefined;
}


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

export async function saveUserAvatarToS3(userResp: UserResp, file: File) {
  if (!userResp.ok) {
    return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  }

  const user_id = userResp.user_id as string

  return { ok: true, resp: await saveToS3("avatar", user_id, file,"profile") }
}

export async function saveUserCoverToS3(userResp: UserResp, file: File) {
  // TODO add some sort of abuse prevention
  if (!userResp.ok) {
    return { ok: userResp.ok, error: userResp.error, message: userResp.message }
  }

  const user_id = userResp.user_id as string

  return { ok: true, resp: await saveToS3("cover", user_id, file, "profile") }
};




type Topic = 'profile' | 'post' | 'group' | 'comment' | 'chat';


export async function saveToS3(type: string, Id: string, file: File, topic: Topic): Promise<string | undefined> {

  const extension = file.name.slice(file.name.lastIndexOf('.'));
  const key = `${topic}/${Id}/${type}${extension}`;

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
    return key
  } catch (error) {
    console.error("S3 upload error", error);
    return
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
}