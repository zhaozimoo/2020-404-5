package main

case "up_load":
upFile, fileHeader, err := req.FormFile("upload_file")
common.PanicError(err, "0020")
fileName := fileHeader.Filename
tmpPath := config.GetCacheDir() + "/" + fileName + uuid.New().String() + ".tmp"
tmpFile, err := os.OpenFile(tmpPath, os.O_RDWR, os.ModePerm)
disk, err := utils.DiskUsage(config.GetCacheDir())
common.PanicError(err, "0021")
if int64(disk.Free) < fileHeader.Size {
common.PanicError(errors.New("no disk spcae"), "0022")
}
total, err := io.Copy(tmpFile, upFile)
if err != nil {
err := utils.DeleteFile(tmpPath)
common.PanicError(err, "0023")
}
common.PanicError(err, "0023")
if total != fileHeader.Size {
common.PanicError(errors.New("upload Failed"), "0024")
}
err = os.Rename(tmpPath, config.GetUploadDir()+"/"+fileName)
common.PanicError(err, "0025")
common.PanicInfo("success", "0000", config.GetUploadDir()+"/"+fileName)