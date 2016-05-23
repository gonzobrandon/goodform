package project3;

import java.io.*;
import java.math.BigDecimal;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.*;
import java.util.stream.Collectors;

/**
 * This class is a data model for the music library. It handles I/O for reading/writing from/to disk. Its independent of GUI.
 *
 * @author Brandon Gonzales
 */
public class MusicLibrary
{
    //column names (not related to keys of Map)
    private String[] keys = {
            "Item Code",
            "Title",
            "Description",
            "Artist",
            "Album",
            "Price"
    };

    private final static String DELIMETER = ";";
    private String filePath;
    private Scanner sc = new Scanner(System.in);
    private boolean inputPause = true;
    private String command;
    private LinkedHashMap<Integer, Track> library = new LinkedHashMap<Integer, Track>();    //Our beloved collection map

    /**
     * This parametrized constructor takes in a file path to read from or save string path for later if its not found.
     *
     * @param filePath (String) the file path of the database file.
     */
    public MusicLibrary(String filePath)
    {

        this.filePath = filePath;

        //If doesn't exist, then tell user and ask if want to create new one.
        if (!new File(filePath).exists())
        {
            System.out.print("\nBad or missing path for music database file.");
            promptForCreateOrExit();
        }
        else
        {
            List<String> stringList = new ArrayList<>();            //Array list for file lines

            try (BufferedReader br = Files.newBufferedReader(Paths.get(filePath)))
            {
                stringList = br.lines().collect(Collectors.toList());       //Stream to list of strings
            }
            catch (IOException e)
            {
                System.out.println("ERROR: There was an I/O error while reading the database file. Exiting.");
                System.exit(0);
            }

            stringList.forEach(this::parseDBLine);      //Lambda functino foreach iterate over lines
        }
    }

    /**
     * This method prompts the user if they want to create a new database file. All CLI prompts.
     *
     */
    public void promptForCreateOrExit()
    {
        inputPause = true;
        while (inputPause)
        {

            System.out.print("\nWould you like to create a new database file? (Y/N):");

            try
            {
                command = sc.nextLine();

                if (command.toLowerCase().equals("y"))
                {
                   //Do nothing here. HashMap was already initialized in declaration.
                }
                else if (command.toLowerCase().equals("n"))
                {
                    System.out.println("GOODBYE!");
                    System.exit(0);
                }
                else
                {
                    throw new IllegalArgumentException();
                }

                inputPause = false;

            }
            catch (IllegalArgumentException e)
            {
                System.out.print("\nError: Please answer Y or N. ");
            }
        }
    }

    /**
     * This method is called by an iterator to build the LinkedHashMap collection line by line.
     *
     * @param line (String) the string of the line read by Bufferedreader.
     */
    private void parseDBLine(String line)
    {

        try
        {
            String[] splitted = line.split(DELIMETER); //Split by delimiter


            //Trim everything
            for (int j = 0; j < splitted.length ; j++)
            {
                splitted[j] = splitted[j].trim();
            }

            //Build new track
            Track track = new Track(splitted[0], splitted[1], splitted[2], splitted[3], new BigDecimal(splitted[4]));

            library.put(System.identityHashCode(track), track);     //Put into collection
        }
        catch (ArrayIndexOutOfBoundsException e)
        {
            System.out.println("ERROR: Improper formatted database music file. Exiting...");
            System.exit(1);
        }
        catch (NumberFormatException e)
        {
            System.out.println("ERROR: Improper formatted database number(s) encountered. Exiting...");
            System.exit(1);
        }

    }

    /**
     * Getters and Setters
     */
    public String[] getKeys()
    {
        return keys;
    }

    public void setKeys(String[] keys)
    {
        this.keys = keys;
    }

    /**
     * This method returns the linked hash map so other classes can access the methods on it.
     *
     * @return (LinkedHashMap<Integer, Track>) The collection of songs.
     */
    public LinkedHashMap<Integer, Track> getLibrary()
    {
        return library;
    }

    /**
     * This method is called when a song is deleted from the UI
     *
     * @param index (int) index of the key to remove
     */
    public void removeAtIndex(int index)
    {
        library.remove(index);
    }

    public void addTrack(Track track)
    {
        library.put(library.size(), track);
    }

    /**
     * This method will save the collectino to disk using the DELIMETER constant. Uses a filewriter/bufferedwriter
     */
    public void saveLibraryToDisk()
    {

        File fileWrite = new File(this.filePath);

        System.out.println("Attempting to Save to Disk to: " + fileWrite);

        try
                (
                        FileWriter fw = new FileWriter(fileWrite);
                        BufferedWriter bw = new BufferedWriter(fw)
                )
        {

            //Iterate over the set via entryset method
            Set<Map.Entry<Integer, Track>> set = library.entrySet();

            for(Map.Entry<Integer, Track> track : set)
            {
                bw.write(track.getValue().toFileRow(DELIMETER) + "\n");
            }

            System.out.println("File saved successfully");
        }
        catch (IOException e)
        {
            System.out.println("ERROR: I/O Error: " + e);
        }
    }

}
