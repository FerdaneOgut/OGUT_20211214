import { useCallback, useEffect, useMemo, useState } from "react";
import { Alert, Button, Container, Form, Stack } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import { Category } from "../data/models/Category";
import { ErrorResponse } from "../data/models/Error";
import CategoryService from "../data/services/categoryService";
import VideoService from "../data/services/videoService";

function VideoUpload(props: any) {
  const navigate = useNavigate();
  const [file, setFile] = useState<any>();
  const [categories, setCategories] = useState<Category[]>([]);
  const [category, setCategory] = useState<string>();
  const [title, setTitle] = useState<string>();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<ErrorResponse>();
  const [success, setSuccess] = useState(false);
  const navigateHome = useCallback(() => {
    navigate("/")
  }, [navigate]);

  const validate = useCallback(() => {
    if (!title) {
      setError({ message: "Title is required" });
      return
    }
    if (!category) {
      setError({ message: "Category is required" });
      return
    }

    if (!file) {
      setError({ message: "File is required" });
      return
    }

  }, [category, title, file])

  const submit = useCallback((e: any) => {
    e.preventDefault();
    //handle validation
    validate();

    setLoading(true);
    const formData = new FormData();
    formData.append('file', file);
    formData.append('title', title!);
    formData.append('category', category!);
    //handle errors;
    VideoService.create(formData)
      .then((response: any) => {
        setSuccess(true);
      })
      .catch((e: any) => {
        const err = e.response.data as ErrorResponse;
        setError(err);
        console.log(err);
      }).finally(() => setLoading(false));
  }, [validate, category, title, file]);

  const onTitleChange = (e: any) => {
    setTitle(e.target.value);
  }
  const onFileChange = (e: any) => {
    const file = e.target.files[0];
    setFile(file);
  }
  const onSelectChange = (e: any) => {
    setCategory(e.target.value);
  };
  useEffect(() => {
    CategoryService.getAll().then(r => setCategories(r.data));
    //handle errors;
  }, [])
  const getOptions = useMemo(() =>
    categories.map((m, i) => <option key={i} value={m.id}>{m.name}</option>)
    , [categories]);

  const onMessageClose = () => {
    setSuccess(false);
    setError(undefined);
  }

  const getMessage = useMemo(() =>
    (error || success) && <Alert style={{ fontSize: "13px" }}
      variant={error ? "danger" : "success"}
      show={(error || success) ? true : false}
      onClose={onMessageClose}
      dismissible>
      {error ? (error.message) : success ? ("File Upload was successfull") : ""}
    </Alert>
    , [error, success]);
  return <Container>

    <Stack style={{ alignItems: "flex-start", marginTop: "10px" }}>
      <Button variant="link" onClick={navigateHome}>Go back</Button>
    </Stack>
    <Stack style={{ alignItems: "center", marginTop: "10px" }}>
      <h3>Upload New Video</h3>
      {getMessage}
      <Form>
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Video Title</Form.Label>
          <Form.Control type="text" onChange={onTitleChange} />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Video Category</Form.Label>
          <Form.Select aria-label="Default select example" onChange={onSelectChange}>
            <option>Select Category</option>
            {getOptions}
          </Form.Select>
        </Form.Group>

        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Label>Video File</Form.Label>
          <Form.Control type="file" onChange={onFileChange} />

        </Form.Group>


        <Button variant="primary" type="submit"
          disabled={loading}
          onClick={!loading ? submit : undefined}
        >
          {loading ? 'Loading…' : 'Click to load'}
        </Button>
      </Form>
    </Stack>
  </Container>

}

export default VideoUpload;