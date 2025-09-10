export default function parseCode(code: string) {
  switch (code) {
    case "FILE_NOT_FOUND":
      return "No file is uploaded";
    case "INTERNAL_SERVICE_ERROR":
      return "Error occured while uploading";
    default:
      return "Something went wrong";
  }
}
