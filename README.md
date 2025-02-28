# PR1

## Как использовать программу

1) Создайте новый проект консольного приложения в Visual Studio

2) Скопируйте код в файл Program.cs

3) Запустите программу

4) Выберите опцию "Открыть файл" и укажите путь к текстовому файлу

5) Введите слово для поиска

6) Программа покажет содержимое файла и количество вхождений слова

```
using System;
using System.IO;
using System.Text.RegularExpressions;

public class FileManager
{
    private readonly string _filePath;

    public FileManager(string filePath)
    {
        _filePath = filePath;
    }

    /// <summary>
    /// Читает содержимое файла
    /// </summary>
    /// <returns>Содержимое файла</returns>
    public string ReadFileContent()
    {
        try
        {
            return File.ReadAllText(_filePath);
        }
        catch (Exception ex)
        {
            throw new Exception($"Ошибка при чтении файла: {ex.Message}");
        }
    }

    /// <summary>
    /// Ищет слово в тексте и возвращает количество вхождений
    /// </summary>
    /// <param name="word">Искомое слово</param>
    /// <returns>Количество вхождений слова</returns>
    public int SearchWord(string word)
    {
        if (string.IsNullOrEmpty(word))
            return 0;

        string content = ReadFileContent();
        
        // Используем регулярные выражения для поиска целых слов
        string pattern = $@"\b{Regex.Escape(word)}\b";
        Regex regex = new Regex(pattern, RegexOptions.IgnoreCase);
        
        return regex.Matches(content).Count;
    }
}

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Консольное приложение для работы с текстовыми файлами");
        Console.WriteLine("------------------------------------------------");

        while (true)
        {
            Console.WriteLine("\nВыберите действие:");
            Console.WriteLine("1. Открыть файл");
            Console.WriteLine("2. Выход");

            string choice = Console.ReadLine();

            switch (choice)
            {
                case "1":
                    OpenAndProcessFile();
                    break;
                case "2":
                    return;
                default:
                    Console.WriteLine("Неверный выбор. Попробуйте снова.");
                    break;
            }
        }
    }

    static void OpenAndProcessFile()
    {
        Console.Write("Введите путь к файлу: ");
        string filePath = Console.ReadLine();

        try
        {
            using (var fileManager = new FileManager(filePath))
            {
                Console.WriteLine("\nСодержимое файла:");
                Console.WriteLine(fileManager.ReadFileContent());

                Console.Write("\nВведите слово для поиска: ");
                string searchWord = Console.ReadLine();

                int count = fileManager.SearchWord(searchWord);
                Console.WriteLine($"\nСлово '{searchWord}' найдено {count} раз.");
            }
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Ошибка: {ex.Message}");
        }

        Console.WriteLine("\nНажмите любую клавишу для продолжения...");
        Console.ReadKey();
    }
}
```
## Особенности реализации:
- Использование регулярных выражений для точного поиска слов с учётом границ слов (\b)
- Работа с исключениями для обработки ошибок при чтении файла
- Автоматическое освобождение ресурсов через using
- Нечувствительность к регистру при поиске (RegexOptions.IgnoreCase)
- Защита от пустых входных данных
