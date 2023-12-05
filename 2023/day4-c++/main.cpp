#include <iostream>
#include <fstream>
#include <istream>
#include <sstream>
#include <vector>
#include <cmath>
#include <numeric>

using namespace std;

std::vector<std::string> split(const std::string &s, char delim)
{
    std::vector<std::string> result;
    std::stringstream ss(s);
    std::string item;

    while (getline(ss, item, delim))
    {
        result.push_back(item);
    }

    return result;
}

void getTotal(vector<string> lines)
{
    int total = 0;
    for (auto i : lines)
    {
        std::vector<std::string> v = split(i, ':');
        std::vector<std::string> sections = split(v[1], '|');
        std::vector<std::string> winning = split(sections[0], ' ');
        std::vector<std::string> numbers = split(sections[1], ' ');
        int count = 0;
        for (auto i : numbers)
        {
            for (auto j : winning)
            {
                if (i == j && i != " " && i != "")
                {
                    count++;
                }
            }
        }
        int points = pow(2, count - 1);
        total += points;
    }
    cout << "Total: " << total << endl;
}

void getScratchCards(vector<string> lines)
{
    int total = 0;
    vector<int> cardCount(lines.size(), 1);

    for (size_t i = 0; i < lines.size(); ++i)
    {
        std::vector<std::string> v = split(lines[i], ':');
        std::vector<std::string> sections = split(v[1], '|');
        std::vector<std::string> winning = split(sections[0], ' ');
        std::vector<std::string> numbers = split(sections[1], ' ');
        int count = 0;
        for (auto i : numbers)
        {
            for (auto j : winning)
            {
                if (i == j && i != " " && i != "")
                {
                    count++;
                }
            }
        }

        for (size_t j = i + 1; j < min(i + 1 + count, lines.size()); ++j)
        {
            cardCount[j] += cardCount[i];
        }
    }

    cout << "Cards: " << accumulate(cardCount.begin(), cardCount.end(), 0) << endl;
}

int main()
{
    string line;

    ifstream InputFile("input.txt");

    int total = 0;
    vector<string> lines;
    while (getline(InputFile, line))
    {
        lines.push_back(line);
    }
    InputFile.close();
    getTotal(lines);
    getScratchCards(lines);
    return 0;
}
