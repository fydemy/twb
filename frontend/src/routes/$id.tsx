import { createFileRoute } from "@tanstack/react-router";
import parseCode from "../lib/code";
import { useQuery } from "@tanstack/react-query";
import type { TwbResponse } from "../lib/types";
import { useDropzone } from "react-dropzone";
import { useCallback, useState } from "react";

const twbQueryOptions = (id: string) => ({
  queryKey: ["twb", id] as const,
  queryFn: async () => {
    const response = await fetch(
      `${import.meta.env.VITE_API_URL}/api/v1/upload/${id}`
    );

    const data = await response.json();
    if (!response.ok) throw Error(parseCode(data.code));

    return data;
  },
  staleTime: 5 * 60 * 1000,
  cacheTime: 10 * 60 * 1000,
  enabled: !!id,
});

export const Route = createFileRoute("/$id")({
  loader: ({ context, params }) =>
    // @ts-expect-error property queryClient does not exist
    context.queryClient.ensureQueryData(twbQueryOptions(params.id)),
  component: RouteComponent,
});

function RouteComponent() {
  const { id } = Route.useParams();
  const { data, isLoading } = useQuery<TwbResponse>(twbQueryOptions(id));
  const [upload, setUpload] = useState<string | null>(null);

  const onDrop = useCallback((acceptedFiles: File[]) => {
    const file = acceptedFiles[0];
    if (!file) return;

    setUpload(URL.createObjectURL(file));
  }, []);

  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    accept: {
      "image/png": [".png"],
      "image/jpeg": [".jpg", ".jpeg"],
    },
    maxFiles: 1,
    multiple: false,
  });

  const handleDownload = async () => {
    if (!data?.path || !upload) return;

    console.log({ data, upload });

    const canvas = document.createElement("canvas");
    const ctx = canvas.getContext("2d");

    const baseImage = new Image();
    const overlayImage = new Image();

    overlayImage.crossOrigin = "anonymous";

    await Promise.all([
      new Promise((resolve) => {
        baseImage.onload = resolve;
        baseImage.src = upload;
      }),
      new Promise((resolve) => {
        overlayImage.onload = resolve;
        overlayImage.src = data.path;
      }),
    ]);

    canvas.width = baseImage.width;
    canvas.height = baseImage.height;

    ctx?.drawImage(baseImage, 0, 0);
    ctx?.drawImage(overlayImage, 0, 0, baseImage.width, baseImage.height);

    const a = document.createElement("a");
    a.href = canvas.toDataURL("image/png");
    a.download = `twibbon-${id}.png`;
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);
  };

  if (isLoading) return "loading..";

  return (
    <>
      <div className="relative aspect-square size-36">
        {upload && (
          <img
            src={upload}
            alt="Base"
            className="absolute inset-0 w-full h-full object-cover"
          />
        )}
        {data?.path && (
          <img
            src={data.path}
            alt="Frame"
            className="absolute inset-0 w-full h-full object-contain z-10"
          />
        )}
      </div>
      <div {...getRootProps()}>
        <input {...getInputProps()} />
        {isDragActive ? (
          <p>Drop the files here ...</p>
        ) : (
          <p>Drag 'n' drop some files here, or click to select files</p>
        )}
      </div>
      {upload && <button onClick={handleDownload}>download</button>}
    </>
  );
}
