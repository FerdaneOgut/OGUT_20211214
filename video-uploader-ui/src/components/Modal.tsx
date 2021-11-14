import { Button, Modal as BModal } from "react-bootstrap"
interface Props {
  show: boolean;
  close: () => void;
  children: any;
}
function Modal(props: Props) {
  const { show, close, children } = props;
  return <BModal show={show} onHide={close} size="lg">
    <BModal.Header closeButton>

    </BModal.Header>
    <BModal.Body>{children}</BModal.Body>
    <BModal.Footer>
      <Button variant="secondary" onClick={close}>
        Close
      </Button>
    </BModal.Footer>
  </BModal>
}

export default Modal;