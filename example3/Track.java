package project3;

import java.math.BigDecimal;

/**
 * This class is a data model for the track object. It containts simple properties for a track.
 *
 * @author Brandon Gonzales
 */
public class Track
{

    private String itemCode = "";
    private String description = "";
    private String artist = "";
    private String album = "";
    private BigDecimal price = new BigDecimal("0");

    /**
     * Parametrized constructor for creating a track. Simple strings and one BigDecimal (could have been string).
     *
     * @param itemCode  (String) The item code
     * @param description (string) the description (title) of the song.
     * @param artist (String) the artist name
     * @param album (String) the album name
     * @param price (BigDecimal) the price of the track.
     */
    public Track(String itemCode,String description, String artist, String album, BigDecimal price)
    {
        this.itemCode = itemCode;
        this.description = description;
        this.artist = artist;
        this.album = album;
        this.price = price;
    }

    /**
     * This constructor is for raw new tracks. Everything is empty or zero.
     */
    public Track()
    {

    }

    /**
     * Getters and setters
     */
    public String getItemCode()
    {
        return itemCode;
    }

    public void setItemCode(String itemCode)
    {
        this.itemCode = itemCode;
    }

    public String getDescription()
    {
        return description;
    }

    public void setDescription(String description)
    {
        this.description = description;
    }

    public String getArtist()
    {
        return artist;
    }

    public void setArtist(String artist)
    {
        this.artist = artist;
    }

    public String getAlbum()
    {
        return album;
    }

    public void setAlbum(String album)
    {
        this.album = album;
    }

    public BigDecimal getPrice()
    {
        return price;
    }

    public void setPrice(BigDecimal price)
    {
        this.price = price;
    }

    /**
     * This method is special for writitng to the database file. Delimeter is argument with space after for readavility
     *
     * @param delimiter (String) The delimeter between data fields. A space is added for readavility
     * @return (String) of the concatenated delimited row
     */
    public String toFileRow(String delimiter)
    {
        return getItemCode() + delimiter + " " + getDescription() + delimiter + " " + getArtist() + delimiter + " " + getAlbum() + delimiter + " " + getPrice().setScale(2, BigDecimal.ROUND_HALF_UP).toString();
    }
}
