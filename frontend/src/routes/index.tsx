import { useCallback } from "react";
import { createFileRoute } from "@tanstack/react-router";
import { useDropzone } from "react-dropzone";
import { useMutation } from "@tanstack/react-query";
import parseCode from "../lib/code";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  const uploadMutation = useMutation({
    mutationFn: async (file: File) => {
      await new Promise<void>((resolve, reject) => {
        const img = new Image();
        img.onload = () => {
          if (img.width !== img.height) {
            reject(new Error("Image must be in square format"));
          } else {
            resolve();
          }
        };
        img.onerror = () => reject(new Error("Error reading image"));
        img.src = URL.createObjectURL(file);
      });

      const form = new FormData();
      form.append("frame", file);

      const response = await fetch(
        `${import.meta.env.VITE_API_URL}/api/v1/upload`,
        {
          method: "POST",
          body: form,
        }
      );

      const data = await response.json();
      if (!response.ok) throw new Error(parseCode(data.code));

      return data;
    },
    onError: (error) => {
      alert(error.message);
    },
    onSuccess: (data) => {
      alert(`${import.meta.env.VITE_APP_URL}/${data.id}`);
    },
  });

  const onDrop = useCallback(
    (acceptedFiles: File[]) => {
      const file = acceptedFiles[0];
      if (!file) return;

      uploadMutation.mutate(file);
    },
    [uploadMutation]
  );

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      "image/png": [".png"],
    },
    maxFiles: 1,
    multiple: false,
  });

  return (
    <div className="p-2">
      <h3>Welcome Home!</h3>
      <div {...getRootProps()}>
        <input {...getInputProps()} />
        {isDragActive ? (
          <p>Drop the files here ...</p>
        ) : (
          <p>Drag 'n' drop some files here, or click to select files</p>
        )}
      </div>
    </div>
  );
}
