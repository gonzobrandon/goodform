package project3;

import javax.swing.*;

/**
 * This class is the primary entry point of the program. It is ran statically and should enter via main().
 * A file location argument is required.
 *
 * * @author Brandon Gonzales 
 */
public class SongDatabase extends JFrame
{

    final static long serialVersionUID = 6292837618161007L;     //For serialization uniqueness
    private static MusicLibrary library;                        //Library passed to JPanel

    /**
     * This consructor is for the JFrame swing extension. This is what is called for the GUI to load.
     */
    public SongDatabase()
    {
        setTitle("JavaTunes");
        setBounds(100, 100, 600, 400);                          //Size
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);         //Kills on close
        setResizable(false);                                    //No resize
        JPanel panel = new SongDatabasePanel(library);             //pass library to the root panel.
        this.add(panel);                                        //Make this panel root panel
    }

    /**
     * This is the main static entry into the program. Ran from the command line.
     *
     * @param args (String[]) the string array of command line args.
     */
    public static void main(String[] args)
    {
        //Check if no args supplied. Exit if empty
        if (args.length < 1)
         {
             System.out.println("You MUST provide the file database file path to read/write from/to as first argument. Exiting....\nGOODBYE!");
             System.exit(1);
         }

        library = new MusicLibrary(args[0]);        //Initialize the
        JFrame frame = new SongDatabase();             //this initializes this class as a JFrame
        frame.setVisible(true);                     //Frame becomes visible
    }

}
