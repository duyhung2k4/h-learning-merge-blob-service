# init process:
  - Tạo ra uuid process => tạo ra 1 queue => hứng blob từ blob_service
  - Call sang stream_service => tạo ra 1 queue => hứng blob từ merge_blob_service 
  - format queue:
    + merge_blob_service: merge_blob_[uuid]
    + stream_service: stream_[uuid]