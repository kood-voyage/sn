import { ListBucketsCommand, PutObjectCommand, S3Client, GetObjectCommand } from '@aws-sdk/client-s3';
import { AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, PROFILE_MEDIA_BUCKET } from '$env/static/private';
import { Readable } from 'stream';


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

export async function saveToS3(path: string, data: object) {
  // TODO add some sort of abuse prevention
  const key = `profile/${path}/banner.json`
  const uploadCommand = new PutObjectCommand({
    Bucket: PROFILE_MEDIA_BUCKET,
    Key: key,
    Body: JSON.stringify(data),
  });
  try {
    const response = await client.send(uploadCommand);
    console.log("S3 upload success ", response);
  } catch (error) {
    console.error("S3 upload error ", error);
  }
};


function streamToString(stream: Readable): Promise<string> {
  return new Promise((resolve, reject) => {
    const chunks: Buffer[] = [];
    stream.on('data', (chunk: Buffer) => chunks.push(chunk));
    stream.once('end', () => resolve(Buffer.concat(chunks).toString('utf-8')));
    stream.on('error', reject);
  });
}

export async function getFromS3(path: string) {
  const key = `/profile/${path}/banner.json`;
  const getCommand = new GetObjectCommand({
    Bucket: PROFILE_MEDIA_BUCKET,
    Key: key,
  });

  try {
    const response = (await client.send(getCommand));
    const Body = response.Body as Readable
    const bodyContents = await streamToString(Body);
    console.log("S3 get success", bodyContents);
    return JSON.parse(bodyContents); // Assuming the stored data is JSON
  } catch (error) {
    console.error("S3 get error", error);
    return null;
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