import { api } from './fetcher';

export interface UploadedFile {
  id: string;
  path: string;
  url: string;
  filename: string;
  contentType: string;
  sizeBytes: number;
}

export interface UploadInput {
  /** URI local do arquivo (ex.: do expo-image-picker / file:// path). */
  uri: string;
  /** Nome do arquivo (preserva extensão para o backend validar). */
  name: string;
  /** MIME type (ex.: image/jpeg). */
  type: string;
}

/**
 * Faz upload de um arquivo via POST /api/v1/uploads (multipart/form-data).
 * Funciona em React Native e web — o caller monta o FormData de acordo.
 */
export async function uploadFile(input: UploadInput): Promise<UploadedFile> {
  const form = new FormData();
  // RN aceita objetos {uri, name, type} no FormData; navegadores usam Blob/File.
  form.append('file', {
    uri: input.uri,
    name: input.name,
    type: input.type,
  } as unknown as Blob);

  return api.post<UploadedFile>('/api/v1/uploads', form);
}
