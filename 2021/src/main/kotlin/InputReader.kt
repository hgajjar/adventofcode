import java.io.File
import java.nio.file.Paths

class InputReader() {
    fun read(fileName: String):  List<String> = File(Paths.get("src/main/input/$fileName").toAbsolutePath().toString()).readLines()
}