import { ListBucketsCommand, PutObjectCommand, S3Client } from '@aws-sdk/client-s3';
import { AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, PROFILE_MEDIA_BUCKET } from '$env/static/private';

const client = new S3Client({
  region: "us-west-2",
  credentials: {
    accessKeyId: AWS_ACCESS_KEY_ID,
    secretAccessKey: AWS_SECRET_ACCESS_KEY
  }
});


const params = {

}

const command = new ListBucketsCommand(params)

// const getKey = (path: string) => `profile/${path}.json`;

export async function saveToS3(path: string, data: object) {
  // TODO add some sort of abuse prevention
  const key = `s3://ImageBucket/profile/${path}/banner`
  const uploadCommand = new PutObjectCommand({
    Bucket: PROFILE_MEDIA_BUCKET,
    Key: key,
    Body: JSON.stringify(data),
    ACL: "public-read",
  });
  try {
    const response = await client.send(uploadCommand);
    console.log("S3 upload success ", response);
  } catch (error) {
    console.error("S3 upload error ", error);
  }
};


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